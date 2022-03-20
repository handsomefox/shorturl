package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/internal/storage"
	"strings"
)

// UnrollRouterPath is the path used for GET method
const UnrollRouterPath = "/u/:link"

// UnrollRouter router itself
type UnrollRouter struct {
	storage *storage.Storage
}

// Interface implementation

// UseStorage points to the Storage for storing the links
func (s *UnrollRouter) UseStorage(storage *storage.Storage) {
	s.storage = storage
}

// Get redirects you to the original link
func (s *UnrollRouter) Get(c *gin.Context) {
	link := c.Param("link")

	found, err := s.storage.Get(link)
	if err != nil {
		return
	}

	if !strings.Contains(found, "http://") {
		if !strings.Contains(found, "https://") {
			found = "https://" + found
		}
	}
	c.Redirect(http.StatusPermanentRedirect, found)
	c.Abort()
}
