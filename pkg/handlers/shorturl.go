package handlers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"shorturl/pkg/encoding"
)

// GetShortenedPath is the path used for GET method
const GetShortenedPath = "/s/:link"

// GetShortened returns the shortened link
func (s *handler) GetShortened(c *gin.Context) {
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
	responseLink, err := s.create(c, decoded)
	log.Printf("Sending Reponse: %s\n", responseLink)
	if err != nil {
		handleError(c, http.StatusBadGateway)
		return
	}

	// Send the link.
	s.sendResponse(c, responseLink)
}

// GetUnshortenedPath is the path used for GET method
const GetUnshortenedPath = "/u/:link"

// GetUnshortened redirects you to the original link
func (s *handler) GetUnshortened(c *gin.Context) {
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
