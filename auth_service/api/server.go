package api

import (
	db "github.com/AbdulRehman-z/instagram-microservices/auth_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/token"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/util"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/worker"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	Config          util.Config
	store           db.Store
	rStore          *redis.Client
	router          *fiber.App
	tokenMaker      token.TokenMaker
	taskDistributor worker.Distributor
}

func NewServer(config util.Config, db db.Store, redisClient *redis.Client, taskDistributor worker.Distributor) (*Server, error) {
	app := fiber.New()
	tokenMaker, err := token.NewPaestoMaker(config.SYMMETRIC_KEY)
	if err != nil {
		return nil, err
	}

	server := &Server{
		Config:          config,
		store:           db,
		rStore:          redisClient,
		router:          app,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
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
	server.router.Get("/health", server.Health)
	server.router.Post("/signup", server.RegisterUser)
	server.router.Post("/login", server.LoginUser)
	// server.router.Post("/login", server.login)
	// server.router.Post("/verify-email", server.verifyEmail)
	// server.router.Post("/refresh", server.refresh)
	// server.router.Post("/logout", server.logout)
	// server.router.Post("/forgot-password", server.forgotPassword)
	// server.router.Post("/reset-password", server.resetPassword)
}
