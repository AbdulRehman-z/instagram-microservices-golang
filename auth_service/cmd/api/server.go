package api

import (
	db "github.com/AbdulRehman-z/instagram-microservices/auth_service/cmd/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/cmd/util"
	"github.com/gofiber/fiber"
)

type Server struct {
	Config util.Config
	db     db.Store
	router *fiber.App
}

func NewServer(config util.Config, db db.Store) (*Server, error) {
	app := fiber.New()
	server := &Server{
		Config: config,
		db:     db,
		router: app,
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
