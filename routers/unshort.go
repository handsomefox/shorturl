package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/storage"
	"strings"
)

const UnshortRouterPath = "/unshort/:link"

// UnshortRouter interface represents the functions which Short struct has
type UnshortRouter interface {
	UseStorage(*storage.Storage)
	Get(c *gin.Context)
}

type Unshort struct {
	storage *storage.Storage
}

func (s *Unshort) UseStorage(storage *storage.Storage) {
	s.storage = storage
}

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
	//data, _ := json.Marshal(found)
	//c.IndentedJSON(http.StatusOK, string(data))
	c.Redirect(http.StatusPermanentRedirect, found)
	c.Abort()
}
