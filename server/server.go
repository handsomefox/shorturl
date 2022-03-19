package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"shorturl/routers"
	"shorturl/storage"
)

// Serv is an interface that describes available APIs for the Server struct
type Serv interface {
	AddGet(string, gin.HandlerFunc)
	AddPost(string, gin.HandlerFunc)
	AddPut(string, gin.HandlerFunc)
	AddDelete(string, gin.HandlerFunc)
}

// Server struct is the object required to start the application
type Server struct {
	Address       string
	engine        *gin.Engine
	storage       *storage.Storage
	shortRouter   *routers.Short
	unshortRouter *routers.Unshort
}

// Run does the setup and launches the server ready to do what it needs to do
func (s *Server) Run() error {
	s.doSetup()
	return s.engine.Run(s.Address)
}

// doSetup adds required routers and APIs before launching the server
func (s *Server) doSetup() {
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
	s.shortRouter = &routers.Short{}
	s.shortRouter.UseStorage(s.storage)
	s.AddGet(routers.ShortRouterPath, s.shortRouter.Get)

	s.unshortRouter = &routers.Unshort{}
	s.unshortRouter.UseStorage(s.storage)
	s.AddGet(routers.UnshortRouterPath, s.unshortRouter.Get)

}

// AddGet Adds a get handler for a given link {path}
func (s *Server) AddGet(path string, f gin.HandlerFunc) {
	s.engine.GET(path, f)
}

// AddPost Adds a post handler for a given link {path}
func (s *Server) AddPost(path string, f gin.HandlerFunc) {
	s.engine.POST(path, f)
}

// AddPut Adds a put handler for a given link {path}
func (s *Server) AddPut(path string, f gin.HandlerFunc) {
	s.engine.PUT(path, f)
}

// AddDelete Adds a delete handler for a given link {path}
func (s *Server) AddDelete(path string, f gin.HandlerFunc) {
	s.engine.DELETE(path, f)
}
