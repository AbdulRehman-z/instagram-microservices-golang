package main

import (
	"database/sql"
	"log/slog"

	"github.com/AbdulRehman-z/instagram-microservices/auth_service/api"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/cache"
	db "github.com/AbdulRehman-z/instagram-microservices/auth_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/util"
	_ "github.com/lib/pq"
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
	util.RunMigration(config.DB_MIGRATION_URL, config.DB_URL)

	// redis client
	redisClient := cache.NewRedisClient(config.REDIS_HOST, config.REDIS_PORT, config.REDIS_PASSWORD, 0)

	store := db.NewStore(conn)
	server, err := api.NewServer(*config, store, redisClient)
	if err != nil {
		slog.Error("Cannot create server: ", err)
	}

	err = server.Start(config.LISTEN_ADDR)
	if err != nil {
		slog.Error("Failed to start server: ", err)
	}
}
