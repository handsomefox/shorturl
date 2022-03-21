package server

import (
	"net/http"
	"shorturl/internal/storage"
	"strings"

	"github.com/gin-gonic/gin"
)

// UnrollRouterPath is the path used for GET method
const UnrollRouterPath = "/u/:link"

// UnrollRouter router itself
type UnrollRouter struct {
	storage storage.Database
}

// Interface implementation

// UseStorage points to the File for storing the links
func (s *UnrollRouter) UseStorage(storage storage.Database) {
	s.storage = storage
}

// Get redirects you to the original link
func (s *UnrollRouter) Get(c *gin.Context) {
	link := c.Param("link")

	found, err := s.storage.Get(link)
	if err != nil {
		return
	}

	if !strings.HasPrefix(found.URL, "http://") {
		if !strings.HasPrefix(found.URL, "https://") {
			found.URL = "https://" + found.URL
		}
	}

	if strings.HasPrefix(found.URL, "localhost:") {
		found.URL = "http://" + found.URL
	}

	c.Redirect(http.StatusPermanentRedirect, found.URL)
	c.Abort()
}
