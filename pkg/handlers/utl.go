package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"net/url"
	"shorturl/pkg/encoding"
	"shorturl/pkg/shortener"
	"shorturl/pkg/storage"
)

func (s *handler) create(c *gin.Context, URL string) (string, error) {
	// Generate a hash from URL.
	hash := shortener.Make(URL)

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

func (s *handler) appendHostname(c *gin.Context, hash string) string {
	return c.Request.Host + "/u/" + hash
}

func (s *handler) sendResponse(c *gin.Context, full string) {
	data := fmt.Sprintf("{\"link\": \"%s\"}", full)
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(data))
}

func isDuplicate(s storage.LinkStorage, link string) bool {
	return s.Contains(link)
}

func getExistingShortLink(s storage.LinkStorage, link string) (string, error) {
	ent, err := s.Get(bson.M{"url": link})
	return ent.Short, err
}

func isValidURL(link string) bool {
	u, err := url.Parse(link)
	return err == nil && (u.Host != "" || u.Path != "")
}

func handleError(c *gin.Context, HTTPStatus int) {
	c.AbortWithStatus(HTTPStatus)
	c.Done()
}
