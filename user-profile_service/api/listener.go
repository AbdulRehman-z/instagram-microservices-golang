package api

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func (s *Server) Listener() error {
	fmt.Println("|||||||||||||------- USER-PROFILE-SERVICE LISTENER STARTED! -------||||||||||||||")
	var (
		FOLLOWERS_FOLLOWING_COUNT_STREAM = "followers_following_count_stream"
		FOLLOWERS_FOLLOWING_GROUP        = "followers_following_group"
		FOLLOWERS_FOLLOWING_CONSUMER_1   = "followers_following_consumer_1"
	)

	_, err := s.redisClient.XGroupCreate(context.Background(), FOLLOWERS_FOLLOWING_COUNT_STREAM,
		FOLLOWERS_FOLLOWING_GROUP, "0").Result()
	if err.Error() == "BUSYGROUP Consumer Group name already exists" {
		log.Printf("group already exists: %s", FOLLOWERS_FOLLOWING_GROUP)
	} else {
		return fmt.Errorf("err creating group: %d", err)
	}

	for {
		streams, err := s.redisClient.XReadGroup(context.Background(), &redis.XReadGroupArgs{
			Streams:  []string{FOLLOWERS_FOLLOWING_COUNT_STREAM, ">"},
			Group:    FOLLOWERS_FOLLOWING_GROUP,
			Consumer: FOLLOWERS_FOLLOWING_CONSUMER_1,
			Block:    0,
			NoAck:    false,
		}).Result()
		if err != nil {
			fmt.Printf("err creating group: %d", err)
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				fmt.Println("message: ", message)
			}
		}
	}
}
