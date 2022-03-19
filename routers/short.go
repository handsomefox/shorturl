package routers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/shortener"
	"shorturl/storage"
)

const ShortRouterPath = "/short/:link"

// ShortRouter interface represents the functions which Short struct has
type ShortRouter interface {
	UseStorage(*storage.Storage)
	Get(c *gin.Context)
}

// Short router itself
type Short struct {
	storage *storage.Storage
}

func (s *Short) UseStorage(storage *storage.Storage) {
	s.storage = storage
}

func (s *Short) Get(c *gin.Context) {
	link := c.Param("link")

	short, full := shortener.Make(link)

	data, _ := json.Marshal(full)
	c.IndentedJSON(http.StatusOK, string(data))

	s.storage.Store(link, short)
}
