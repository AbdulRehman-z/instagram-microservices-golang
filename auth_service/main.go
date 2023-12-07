package main

import (
	"database/sql"
	"log/slog"

	"github.com/AbdulRehman-z/instagram-microservices/auth_service/api"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/cache"
	db "github.com/AbdulRehman-z/instagram-microservices/auth_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/mail"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/util"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/worker"
	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
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
	util.RunMigration(config.DB_MIGRATION_URL, config.DB_URL)

	// redis client
	redisClient := cache.NewRedisClient(config.REDIS_HOST, config.REDIS_PORT, config.REDIS_PASSWORD, 0)

	// task distributor
	options := asynq.RedisClientOpt{
		Addr: redisClient.Options().Addr,
	}
	distributor := worker.NewDistributor(&options)

	store := db.NewStore(conn)
	go run(config, store, redisClient, distributor)

}

func run(config *util.Config, store db.Store, redisClient *redis.Client, distributor worker.Distributor) {
	server, err := api.NewServer(*config, store, redisClient, distributor)
	if err != nil {
		slog.Error("Cannot create server: ", err)
	}

	err = server.Start(config.LISTEN_ADDR)
	if err != nil {
		slog.Error("Failed to start server: ", err)
	}
}

func runTaskProcessor(options asynq.RedisClientOpt, store db.Store, mail mail.Mailer) {
	processor := worker.NewProcessor(&options, store, mail)
	err := processor.Start()
	if err != nil {
		slog.Error("failed to start task processor", "err", err)
	}
}
