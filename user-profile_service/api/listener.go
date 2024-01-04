package api

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func (s *Server) FollowersListener() error {
	fmt.Println("|||||||||||||------- FOLLOWERS LISTENER STARTED! -------||||||||||||||")
	var (
		FOLLOWERS_FOLLOWING_COUNT_STREAM = "followers_following_count_stream"
	)

	for {
		streams, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{FOLLOWERS_FOLLOWING_COUNT_STREAM, "$"},
			Block:   0,
		}).Result()
		if err != nil {
			fmt.Printf("err creating : %d", err)
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				fmt.Println("message values: \n", message.Values)
			}
		}
	}
}

func (s *Server) PostsListener() error {
	fmt.Println("|||||||||||||------- POSTS LISTENER STARTED! -------||||||||||||||")
	var (
		POSTS_STREAM = "posts_stream"
	)

	for {
		streams, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{POSTS_STREAM, "$"},
			Block:   0,
		}).Result()
		if err != nil {
			log.Printf("err reading from stream:%s || err: %s\n", POSTS_STREAM, err)
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				log.Println("message values: \n", message.Values)
			}
		}
	}
}

func (s *Server) AccountListener() {
	fmt.Println("|||||||||||||------- ACCOUNTS LISTENER STARTED! -------||||||||||||||")
	var (
		ACCOUNT_STREAM = "account_stream"
	)

	for {
		streams, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{ACCOUNT_STREAM, "$"},
			Block:   0,
		}).Result()
		if err != nil {
			log.Printf("err reading from stream:%s || err: %s\n", ACCOUNT_STREAM, err)
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				log.Println("message values: \n", message.Values)
			}
		}
	}
}
