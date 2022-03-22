package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"shorturl/internal/storage"
	"shorturl/pkg/shortener"
	"strings"
)

// ShortRouterPath is the path used for GET method
const ShortRouterPath = "/s/:link"

// ShortRouter - router itself
type ShortRouter struct {
	storage storage.Database
}

// Interface implementation

// UseStorage points to the File for storing the links
func (s *ShortRouter) UseStorage(storage storage.Database) {
	s.storage = storage
}

// Get returns the shortened link
func (s *ShortRouter) Get(c *gin.Context) {
	str := c.Param("link")
	_, err := url.Parse(str)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	str = strings.ReplaceAll(str, "|", "/")
	str = strings.ReplaceAll(str, "{", "?")
	str = strings.ReplaceAll(str, "}", "=")
	str = strings.ReplaceAll(str, "[", "&")

	short, full := shortener.Make(c.Request.Host, str)

	data := fmt.Sprintf("{\"link\": \"%s\"}", full)

	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(data))
	s.storage.Store(str, short)
}
