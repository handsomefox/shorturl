package server

import (
	"os"
	"shorturl/pkg/handlers"
	"shorturl/pkg/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// MongoServer struct is the object required to start the application
type MongoServer struct {
	Address string
	DBKey   string
	engine  *gin.Engine
	storage storage.LinkStorage
}

// New returns a new MongoServer with given parameters
func New(address string, dbKey string) *MongoServer {
	gin.SetMode(gin.ReleaseMode)
	s := &MongoServer{
		Address: address,
		DBKey:   dbKey,
		engine:  gin.Default(),
		storage: &storage.Database{Key: dbKey},
	}
	s.engine.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

	viewPath := os.Getenv("VIEW_PATH")
	s.engine.Use(static.Serve("/", static.LocalFile(viewPath, true)))
	return s
}

// Run launches the server and catches errors
func (s *MongoServer) Run() error {
	return s.engine.Run(s.Address)
}

// Init adds required routers and initializes the storage
func (s *MongoServer) Init() error {
	if err := s.storage.Init(); err != nil {
		return err
	}

	h := handlers.New(s.storage)
	s.engine.GET(handlers.GetShortenedPath, h.GetShortened)
	s.engine.GET(handlers.GetUnshortenedPath, h.GetUnshortened)

	return nil
}
