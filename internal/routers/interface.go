package routers

import (
	"github.com/gin-gonic/gin"
	"shorturl/internal/storage"
)

// Route interface describes the functions that routers have, for now
type Route interface {
	UseStorage(*storage.Storage)
	Get(c *gin.Context)
}
