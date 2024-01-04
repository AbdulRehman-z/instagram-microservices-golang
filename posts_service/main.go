package main

import (
	"database/sql"
	"log"

	"github.com/AbdulRehman-z/instagram-microservices/posts_service/api"
	db "github.com/AbdulRehman-z/instagram-microservices/posts_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/posts_service/util"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Printf("Error loading config: %v", err)
	}

	conn, err := sql.Open(config.DB_DRIVER, config.DB_URL)
	if err != nil {
		log.Printf("Cannot connect to DB: %v", err)
	}
	if err := conn.Ping(); err != nil {
		log.Printf("Cannot ping DB: %v", err)
	}

	// util.RunMigration(config.DB_MIGRATION_URL, config.DB_URL)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ADDR,
		Password: config.REDIS_PASSWORD,
		DB:       0,
	})

	store := db.NewStore(conn)
	run(config, store, redisClient)
}

func run(config *util.Config, store db.Store, redisClient *redis.Client) {
	server, err := api.NewServer(*config, store, redisClient)
	if err != nil {
		log.Fatalf("failed to initiate server: %d", err)
	}
	go server.UserProfileListener()
	go server.Publisher()
	err = server.Start(config.LISTEN_ADDR)
	if err != nil {
		log.Fatalf("failed to start server: %d", err)
	}
}
