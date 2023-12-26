package api

import (
	"log"

	db "github.com/AbdulRehman-z/instagram-microservices/followers_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/followers_service/token"
	"github.com/AbdulRehman-z/instagram-microservices/followers_service/util"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Config        *util.Config
	router        *fiber.App
	store         db.Store
	TokenVerifier token.TokenVerifier
}

func NewServer(config *util.Config, store db.Store) (*Server, error) {

	tokenVerifier, err := token.NewPaestoMaker(config.SYMMETRIC_KEY)
	if err != nil {
		log.Fatalf("cannot create token maker: %v", err)
	}

	return &Server{
		Config:        config,
		router:        fiber.New(),
		store:         store,
		TokenVerifier: tokenVerifier,
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
	auth := s.router.Group("/api", AuthMiddleware(s.TokenVerifier))
	auth.Get("/followers", s.GetFollowers)
	auth.Get("/following", s.GetFollowings)
	auth.Post("/follow", s.FollowUser)
	auth.Delete("/unfollow", s.UnfollowUser)
	auth.Get("/following/count", s.GetFollowingsCount)
	auth.Get("/followers/count", s.GetFollowersCount)
}
