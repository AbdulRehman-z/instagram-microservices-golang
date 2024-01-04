package main

import (
	"database/sql"
	"log"

	"github.com/AbdulRehman-z/instagram-microservices/likes_service/api"
	db "github.com/AbdulRehman-z/instagram-microservices/likes_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/likes_service/util"
	"github.com/redis/go-redis/v9"
)

func testDb() *sql.DB {
	return &sql.DB{}
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("err loading config: %d", err)
	}

	conn := testDb()
	store := db.NewStore(conn)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ADDR,
		Password: config.REDIS_PASSWORD,
		DB:       0,
	})

	run(config, store, redisClient)
}

func run(config *util.Config, store db.Store, redisClient *redis.Client) {
	server, err := api.NewServer(*config, store, redisClient)
	if err != nil {
		log.Fatal(err)
	}
	go server.CommentsLikesListener()
	go server.PostsLikesListener()
	go server.CommentsLikesPublisher()
	go server.PostsLikesPublisher()
	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
