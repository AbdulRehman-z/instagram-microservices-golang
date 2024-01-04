package api

import (
	db "github.com/AbdulRehman-z/instagram-microservices/create-account_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/create-account_service/token"
	"github.com/AbdulRehman-z/instagram-microservices/create-account_service/util"
	"github.com/redis/go-redis/v9"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	Config        util.Config
	store         db.Store
	router        *fiber.App
	tokenVerifier token.TokenVerifier
	redisClient   *redis.Client
	uniqueId      string
	accountChan   chan *Account
}

func NewServer(config util.Config, db db.Store, redisClient *redis.Client) (*Server, error) {
	app := fiber.New()
	tokenVerifier, err := token.NewPaestoMaker(config.SYMMETRIC_KEY)
	if err != nil {
		return nil, err
	}

	app.Use(logger.New(logger.ConfigDefault))
	server := &Server{
		Config:        config,
		store:         db,
		router:        app,
		tokenVerifier: tokenVerifier,
		redisClient:   redisClient,
		uniqueId:      "550e8400-e29b-41d4-a716-446655440000",
		accountChan:   make(chan *Account),
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

	account := app.Group("/", AuthMiddleware(server.tokenVerifier))
	account.Post("/create_account", server.CreateAccount)
	account.Get("/get_account", server.GetAccount)
	account.Put("/update_account", server.UpdateAccount)
	account.Delete("/delete_account", server.DeleteAccount)
}
