package api

import (
	"github.com/AbdulRehman-z/instagram-microservices/user-profile_service/token"
	"github.com/AbdulRehman-z/instagram-microservices/user-profile_service/util"
	"github.com/redis/go-redis/v9"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	Config        util.Config
	redisClient   redis.Client
	router        *fiber.App
	tokenVerifier token.TokenVerifier
	UniqueId      string
	// totalFollowersChan chan int64
	// totalFollowingChan chan int64
	// totalPostsChan     chan int64
	// postsChan          chan []Post
	// accountChan        chan Account
}

func NewServer(config util.Config, redisClient redis.Client) (*Server, error) {
	app := fiber.New()
	tokenVerifier, err := token.NewPaestoMaker(config.SYMMETRIC_KEY)
	if err != nil {
		return nil, err
	}

	app.Use(logger.New(logger.ConfigDefault))
	server := &Server{
		redisClient:   redisClient,
		Config:        config,
		router:        app,
		tokenVerifier: tokenVerifier,
		UniqueId:      "12345",
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
	server.router.Get("/health", nil)
}
