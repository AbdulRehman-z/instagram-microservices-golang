package main

import (
	"database/sql"
	"log"

	"github.com/AbdulRehman-z/instagram-microservices/comments_service/api"
	db "github.com/AbdulRehman-z/instagram-microservices/comments_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/comments_service/util"
	"github.com/redis/go-redis/v9"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("err loading config: %s", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ADDR,
		Password: config.REDIS_PASSWORD,
		DB:       0,
	})

	run(config, db.NewStore(&sql.DB{}), redisClient)
}

func run(config *util.Config, store db.Store, redisClient *redis.Client) {
	server, err := api.NewServer(config, store, redisClient)
	if err != nil {
		log.Fatalf("err creating server: %s", err)
	}
	go server.Publisher()
	go server.CommentsLikesListener()
	go server.PostsIdsListener()
	err = server.Start()
	if err != nil {
		log.Fatalf("err starting server: %s", err)
	}

}
