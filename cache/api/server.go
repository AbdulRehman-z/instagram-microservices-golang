package api

import (
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type Server struct {
	listenAddr  string
	redisClient *redis.Client
}

func NewServer(listenAddr string, redisClient *redis.Client) *Server {
	return &Server{
		listenAddr:  listenAddr,
		redisClient: redisClient,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))

	log.Printf("Starting server on %s", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, nil)
}
