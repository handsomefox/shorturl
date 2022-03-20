package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"shorturl/internal/routers"
	"shorturl/internal/storage"
)

// Server is an interface that describes available APIs for the ShortURLServer struct
type Server interface {
	AddGET(string, gin.HandlerFunc)
	AddPOST(string, gin.HandlerFunc)
	AddPUT(string, gin.HandlerFunc)
	AddDELETE(string, gin.HandlerFunc)
}

// ShortURLServer struct is the object required to start the application
type ShortURLServer struct {
	Address      string
	engine       *gin.Engine
	storage      *storage.Storage
	routerShort  *routers.ShortRouter
	routerUnroll *routers.UnrollRouter
}

// Run does the setup and launches the server ready to do what it needs to do
func (s *ShortURLServer) Run() error {
	s.doSetup()
	return s.engine.Run(s.Address)
}

// doSetup adds required routers and APIs before launching the server
func (s *ShortURLServer) doSetup() {
	s.engine = gin.Default()

	s.engine.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

	s.storage = &storage.Storage{FilePath: "C:\\Go\\Saved\\data.json"}
	s.storage.Init()

	// Add routers here
	s.routerShort = &routers.ShortRouter{}
	s.routerShort.UseStorage(s.storage)
	s.AddGET(routers.ShortRouterPath, s.routerShort.Get)

	s.routerUnroll = &routers.UnrollRouter{}
	s.routerUnroll.UseStorage(s.storage)
	s.AddGET(routers.UnrollRouterPath, s.routerUnroll.Get)

}

// AddGET Adds a get handler for a given link {path}
func (s *ShortURLServer) AddGET(path string, f gin.HandlerFunc) {
	s.engine.GET(path, f)
}

// AddPOST Adds a post handler for a given link {path}
func (s *ShortURLServer) AddPOST(path string, f gin.HandlerFunc) {
	s.engine.POST(path, f)
}

// AddPUT Adds a put handler for a given link {path}
func (s *ShortURLServer) AddPUT(path string, f gin.HandlerFunc) {
	s.engine.PUT(path, f)
}

// AddDELETE Adds a delete handler for a given link {path}
func (s *ShortURLServer) AddDELETE(path string, f gin.HandlerFunc) {
	s.engine.DELETE(path, f)
}