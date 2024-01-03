package api

import (
	"fmt"

	db "github.com/AbdulRehman-z/instagram-microservices/posts_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/posts_service/token"
	"github.com/AbdulRehman-z/instagram-microservices/posts_service/util"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	Config        util.Config
	store         db.Store
	router        *fiber.App
	TokenVerifier token.TokenVerifier
	redisClient   *redis.Client
	PostsChan     chan *PostEvent
	uniqueId      string
}

func NewServer(config util.Config, store db.Store, redisClient *redis.Client) (*Server, error) {
	paestoVerifier, err := token.NewPaestoMaker(config.SYMMETRIC_KEY)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate paseto verifier: %d", err)
	}

	app := fiber.New()
	server := &Server{
		Config:        config,
		store:         store,
		router:        app,
		TokenVerifier: paestoVerifier,
		redisClient:   redisClient,
		PostsChan:     make(chan *PostEvent),
		uniqueId:      "550e8400-e29b-41d4-a716-446655440000",
	}

	server.SetupRoutes(app)
	return server, nil
}

func (server *Server) Start(listenAddr string) error {
	return server.router.Listen(listenAddr)
}

func (server *Server) Shutdown() error {
	return server.router.Shutdown()
}

func (server *Server) SetupRoutes(app *fiber.App) {
	server.router.Get("/health", server.Health)

	auth := app.Group("/", AuthMiddleware(server.TokenVerifier))
	auth.Post("/posts", server.CreatePost)
	auth.Get("/posts", server.GetPostsByUniqueId)
	auth.Get("/posts/:id", server.GetPostById)
	auth.Put("/posts/:id", server.UpdatePost)
	auth.Delete("/posts/:id", server.DeletePost)
	auth.Delete("/posts", server.DeletePostsByUniqueId)
}
