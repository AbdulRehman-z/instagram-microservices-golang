package api

import (
	db "github.com/AbdulRehman-z/instagram-microservices/auth_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/token"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/util"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	Config      util.Config
	store       db.Store
	redisClient *redis.Client
	router      *fiber.App
	tokenMaker  token.TokenMaker
}

func NewServer(config util.Config, db db.Store, redisClient *redis.Client) (*Server, error) {
	app := fiber.New()
	tokenMaker, err := token.NewPaestoMaker(config.SYMMETRIC_KEY)
	if err != nil {
		return nil, err
	}

	server := &Server{
		Config:      config,
		store:       db,
		redisClient: redisClient,
		router:      app,
		tokenMaker:  tokenMaker,
	}

	server.SetupRoutes()
	return server, nil
}

func (server *Server) Start(listenAddr string) error {
	return server.router.Listen(listenAddr)
}

func (server *Server) Shutdown() error {
	return server.router.Shutdown()
}

func (server *Server) SetupRoutes() {
	// server.router.Get("/health", server.health)
	// server.router.Post("/signup", server.signup)
	// server.router.Post("/login", server.login)
	// server.router.Post("/verify-email", server.verifyEmail)
	// server.router.Post("/refresh", server.refresh)
	// server.router.Post("/logout", server.logout)
	// server.router.Post("/forgot-password", server.forgotPassword)
	// server.router.Post("/reset-password", server.resetPassword)
}
