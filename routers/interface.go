package routers

import (
	"github.com/gin-gonic/gin"
	"shorturl/storage"
)

type Router interface {
	UseStorage(*storage.Storage)
	Get(c *gin.Context)
}
