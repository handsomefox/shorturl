package server

import (
	"shorturl/internal/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
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
	DBKey        string
	engine       *gin.Engine
	storage      storage.Model
	routerShort  Route
	routerUnroll Route
}

// Run does the setup and launches the server ready to do what it needs to do
func (s *ShortURLServer) Run() error {
	return s.engine.Run(s.Address)
}

// Init adds required routers and APIs before launching the server
func (s *ShortURLServer) Init() error {
	gin.SetMode(gin.ReleaseMode)
	s.engine = gin.Default()

	s.engine.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

	s.engine.Use(static.Serve("/", static.LocalFile("./views", true)))

	s.storage = &storage.Database{Key: s.DBKey}
	err := s.storage.Init()

	if err != nil {
		return err
	}

	// Add routers here
	s.routerShort = &ShortRouter{storage: s.storage}
	s.AddGET(ShortRouterPath, s.routerShort.Get)

	s.routerUnroll = &UnrollRouter{storage: s.storage}
	s.AddGET(UnrollRouterPath, s.routerUnroll.Get)
	return nil
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
