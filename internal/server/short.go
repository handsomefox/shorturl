package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"shorturl/internal/storage"
	"shorturl/pkg/shortener"
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
	parse, err := url.Parse(c.Param("link"))

	if err != nil {
		log.Fatal(err)
		return
	}

	short, full := shortener.Make(c.Request.Host, parse.String())

	json := fmt.Sprintf("{\"link\": \"%s\"}", full)

	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(json))
	s.storage.Store(parse.String(), short)
}
