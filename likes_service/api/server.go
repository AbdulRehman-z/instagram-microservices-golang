package api

import (
	db "github.com/AbdulRehman-z/instagram-microservices/likes_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/likes_service/token"
	"github.com/AbdulRehman-z/instagram-microservices/likes_service/util"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Config        util.Config
	store         db.Store
	TokenVerifier token.TokenVerifier
	router        *fiber.App
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenVerifier, err := token.NewPaestoMaker(config.SYMMETRIC_KEY)
	if err != nil {
		return nil, err
	}

	app := fiber.New()

	return &Server{
		Config:        config,
		store:         store,
		router:        app,
		TokenVerifier: tokenVerifier,
	}, nil
}

func (s *Server) Start() error {
	s.Routes()
	return s.router.Listen(s.Config.LISTEN_ADDR)
}

func (s *Server) Shutdown() error {
	return s.router.Shutdown()
}

func (s *Server) Routes() {
	s.router.Get("/health", s.HealthCheck)

	// AuthMiddleware is a middleware that checks if the request is authorized
	auth := s.router.Group("/api", AuthMiddleware(s.TokenVerifier))

	// Post likes api endpoints
	auth.Post("/like/post", s.LikePost)
	auth.Delete("/like/post", s.UnlikePost)
	auth.Get("/like/post/count", s.GetPostLikesCount)
	auth.Get("/like/post", s.GetPostLikes)

	// Comment likes api endpoints
	auth.Post("/like/commment", s.LikeComment)
	auth.Delete("/like/comment", s.UnlikeComment)
	auth.Get("/like/comment/count", s.GetPostLikesCount)
}
