package routers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/shortener"
	"shorturl/storage"
)

// ShortRouterPath is the path used for GET method
const ShortRouterPath = "/short/:link"

// Short router itself
type Short struct {
	storage *storage.Storage
}

// Interface implementation

// UseStorage points to the Storage for storing the links
func (s *Short) UseStorage(storage *storage.Storage) {
	s.storage = storage
}

// Get returns the shortened link
func (s *Short) Get(c *gin.Context) {
	link := c.Param("link")

	short, full := shortener.Make(link)

	data, _ := json.Marshal(full)
	c.IndentedJSON(http.StatusOK, string(data))

	s.storage.Store(link, short)
}
