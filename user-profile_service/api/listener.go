package api

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (s *Server) Listener() {
	for {
		fmt.Println("||||||||||||| LISTENER ||||||||||||||")

		// streams, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
		// 	Streams: []string{"followers_following_count", "0"}, // "0" means "read from the beginning
		// 	Block:   0,
		// }).Result()
		// if err != nil {
		// 	fmt.Printf("error reading from stream: %d", err)
		// }

		res, err := s.redisClient.XGroupCreate(context.Background(), "followers_following_count",
			"followers_following_group", "0").Result()
		if err != nil {
			fmt.Printf("err creating group: %d", err)
		}
		fmt.Println(res)

		streams, err := s.redisClient.XReadGroup(context.Background(), &redis.XReadGroupArgs{
			Streams: []string{"followers_following_count", ">"},
			Group:   "followers_following_group",
			Block:   0,
			NoAck:   false,
		}).Result()
		if err != nil {
			fmt.Printf("err creating group: %d", err)
		}

		fmt.Println("streams: ", streams)
		for _, stream := range streams {
			for _, message := range stream.Messages {
				fmt.Println("message: ", message)
			}
		}
	}
}
