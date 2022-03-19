package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/storage"
	"strings"
)

// UnshortRouterPath is the path used for GET method
const UnshortRouterPath = "/unshort/:link"

// Unshort router itself
type Unshort struct {
	storage *storage.Storage
}

// Interface implementation

// UseStorage points to the Storage for storing the links
func (s *Unshort) UseStorage(storage *storage.Storage) {
	s.storage = storage
}

// Get redirects you to the original link
func (s *Unshort) Get(c *gin.Context) {
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
