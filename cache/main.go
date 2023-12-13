package main

import (
	"log"

	"github.com/instagram-microservices/cache/api"
	"github.com/instagram-microservices/cache/redis"
	"github.com/instagram-microservices/cache/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	redisClient := redis.NewRedisClient(config.REDIS_HOST, config.REDIS_PORT, config.REDIS_PASSWORD, 0)
	server := api.NewServer(config.LISTEN_ADDR, redisClient.Client)

	server.Start()
}
