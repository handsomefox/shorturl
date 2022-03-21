package server

import (
	"shorturl/internal/storage"

	"github.com/gin-gonic/gin"
)

// Route interface describes the functions that routers have, for now
type Route interface {
	UseStorage(storage.Model)
	Get(c *gin.Context)
}
