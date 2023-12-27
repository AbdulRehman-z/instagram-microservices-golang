package api

import (
	db "github.com/AbdulRehman-z/instagram-microservices/comments_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/comments_service/token"
	"github.com/AbdulRehman-z/instagram-microservices/comments_service/util"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Config        util.Config
	router        *fiber.App
	store         db.Store
	tokenVerifier token.TokenVerifier
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	tokenVerifier, err := token.NewPaestoMaker(config.SYMMETRIC_KEY)
	if err != nil {
		return nil, err
	}

	return &Server{
		Config:        config,
		router:        fiber.New(),
		store:         store,
		tokenVerifier: tokenVerifier,
	}, nil
}

func (s *Server) Start() error {
	return s.router.Listen(s.Config.LISTEN_ADDR)
}

func (s *Server) Shutdown() error {
	return s.router.Shutdown()
}

func (s *Server) Routes() {
	s.router.Get("/health", s.HealthCheck)

	// AuthMiddleware is a middleware that checks if the request is authorized
	auth := s.router.Group("/api", AuthMiddleware(s.tokenVerifier))
	auth.Post("/comment", s.CreateComment)
	auth.Get("/comments", s.GetComments)
	auth.Get("/comments/count", s.GetCommentsCount)
	auth.Put("/comment", s.UpdateComment)
	auth.Delete("/comment", s.DeleteComment)
}
