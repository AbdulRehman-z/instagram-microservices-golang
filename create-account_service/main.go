package main

import (
	"database/sql"
	"log/slog"

	"github.com/AbdulRehman-z/instagram-microservices/create-account_service/api"
	db "github.com/AbdulRehman-z/instagram-microservices/create-account_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/create-account_service/util"
	"github.com/redis/go-redis/v9"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		slog.Error("Error loading config: ", err)
	}

	conn, err := sql.Open(config.DB_DRIVER, config.DB_URL)
	if err != nil {
		slog.Error("Cannot connect to DB: ", err)
	}

	err = conn.Ping()
	if err != nil {
		slog.Error("Cannot ping DB: ", err)
	}

	defer conn.Close()
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
		slog.Error("Cannot create server: ", err)
	}
	go server.Publisher()
	go server.UserProfileListener()
	err = server.Start(config.LISTEN_ADDR)
	if err != nil {
		slog.Error("Failed to start server: ", err)
	}
}
