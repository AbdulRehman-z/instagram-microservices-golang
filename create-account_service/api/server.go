package api

import (
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/util"
	db "github.com/AbdulRehman-z/instagram-microservices/create-account_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/create-account_service/token"

	// "github.com/AbdulRehman-z/instagram-microservices/create-account_service/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	Config        util.Config
	store         db.Store
	rStore        *redis.Client
	router        *fiber.App
	tokenVerifier token.TokenVerifier
}

func NewServer(config util.Config, db db.Store, redisClient *redis.Client) (*Server, error) {
	app := fiber.New()
	tokenMaker, err := token.NewPaestoMaker(config.SYMMETRIC_KEY)
	if err != nil {
		return nil, err
	}

	app.Use(logger.New(logger.ConfigDefault))
	server := &Server{
		Config:        config,
		store:         db,
		rStore:        redisClient,
		router:        app,
		tokenVerifier: tokenMaker,
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
	account.Post("/create_account", nil)
	account.Get("/get_account", nil)
	account.Put("/update_account", nil)
	account.Delete("/delete_account", nil)
}
