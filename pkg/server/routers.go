package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"net/url"
	"shorturl/pkg/encoding"
	"shorturl/pkg/shorturl"
	"shorturl/pkg/storage"
)

// Route interface describes the functions that routers have, for now
type Route interface {
	Get(c *gin.Context)
}

// ShortRouterPath is the path used for GET method
const ShortRouterPath = "/s/:link"

// UnrollRouterPath is the path used for GET method
const UnrollRouterPath = "/u/:link"

// ShortRouter - router itself
type ShortRouter struct {
	storage storage.Model
}

// UnrollRouter router itself
type UnrollRouter struct {
	storage storage.Model
}

// Interface implementation

// Get returns the shortened link
func (s *ShortRouter) Get(c *gin.Context) {
	// Receive a base64 encoded URL.
	encoded := c.Param("link")
	// Decode it.
	decoder := encoding.NewDecoder()
	decoded, err := decoder.Decode(encoded)

	log.Printf("Recieved URL: %s\n", decoded)

	if err != nil {
		handleError(c, http.StatusInternalServerError)
		return
	}

	// Check if URL is valid.
	if !isValidURL(decoded) {
		handleError(c, http.StatusBadRequest)
		return
	}

	// Check if we already store that URL.
	if isDuplicate(s.storage, decoded) {
		dup, err := getExistingShortLink(s.storage, decoded)
		if err != nil {
			handleError(c, http.StatusInternalServerError)
			return
		}
		// Return existing URL instead.
		s.sendResponse(c, s.appendHostname(c, dup))
		return
	}

	// Create a short link and store it in the database.
	responseLink, err := s.makeAndStore(c, decoded)
	log.Printf("Sending Reponse: %s\n", responseLink)
	if err != nil {
		handleError(c, http.StatusBadGateway)
		return
	}

	// Send the link.
	s.sendResponse(c, responseLink)
}

func (s *ShortRouter) makeAndStore(c *gin.Context, URL string) (string, error) {
	// Generate a hash from URL.
	hash := shorturl.Make(URL)

	// Store the hash.
	err := s.storage.Store(hash, URL)

	log.Printf("Stored the hash: %s\n", hash)

	// Encode it for the link.
	encoder := encoding.NewEncoder()
	encoded := encoder.Encode(hash)

	// Add a hostname so the hash can be used as a link.
	full := s.appendHostname(c, encoded)
	return full, err
}

func (s *ShortRouter) appendHostname(c *gin.Context, hash string) string {
	return c.Request.Host + "/u/" + hash
}

func (s *ShortRouter) sendResponse(c *gin.Context, full string) {
	data := fmt.Sprintf("{\"link\": \"%s\"}", full)
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(data))
}

// Get redirects you to the original link
func (s *UnrollRouter) Get(c *gin.Context) {
	// Get a base64 encoded hash.
	b64 := c.Param("link")

	// Decode the hash.
	decoder := encoding.NewDecoder()
	hash, err := decoder.Decode(b64)
	if err != nil {
		handleError(c, http.StatusInternalServerError)
		return
	}

	// Search the hash in database.
	found, err := s.storage.Get(bson.M{"short": hash})
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Return the found URL.
	c.Redirect(http.StatusPermanentRedirect, found.URL)
	c.Abort()
}

func handleError(c *gin.Context, HTTPStatus int) {
	c.AbortWithStatus(HTTPStatus)
	c.Done()
}

func isDuplicate(s storage.Model, link string) bool {
	return s.Contains(link)
}

func getExistingShortLink(s storage.Model, link string) (string, error) {
	ent, err := s.Get(bson.M{"url": link})
	return ent.Short, err
}

func isValidURL(link string) bool {
	u, err := url.Parse(link)
	return err == nil && (u.Host != "" || u.Path != "")
}
