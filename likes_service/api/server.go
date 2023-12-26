package api

import (
	db "github.com/AbdulRehman-z/instagram-microservices/likes_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/likes_service/util"
	"github.com/AbdulRehman-z/instagram-microservices/posts_service/token"
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
	auth.Post("/likes/post", s.LikePost)
	auth.Delete("/likes/post", s.UnlikePost)
	auth.Get("/likes/post/count", s.GetPostLikesCount)
	auth.Get("/likes/post", s.GetPostLikes)

	// Comment likes api endpoints
	auth.Post("/likes/commment", s.LikeComment)
	auth.Delete("/likes/comment", s.UnlikeComment)
	auth.Get("/likes/comment/count", s.GetPostLikesCount)
}
