package routers

import (
	"github.com/gin-gonic/gin"
	"shorturl/storage"
)

// Router interface describes the functions that routers have, for now
type Router interface {
	UseStorage(*storage.Storage)
	Get(c *gin.Context)
}
