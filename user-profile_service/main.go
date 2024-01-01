package main

import (
	"log/slog"

	"github.com/AbdulRehman-z/instagram-microservices/user-profile_service/api"
	"github.com/AbdulRehman-z/instagram-microservices/user-profile_service/util"
	"github.com/redis/go-redis/v9"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		slog.Error("Error loading config: ", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ADDR,
		Password: config.REDIS_PASSWORD,
		DB:       0,
	})

	run(config, *redisClient)
}

func run(config *util.Config, redisClient redis.Client) {
	server, err := api.NewServer(*config, redisClient)
	if err != nil {
		slog.Error("Cannot create server: ", err)
	}
	go server.Listener()
	go server.Publisher(server.UniqueId)
	err = server.Start(config.LISTEN_ADDR)
	if err != nil {
		slog.Error("Failed to start server: ", err)
	}
}
