package api

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func (s *Server) UserProfileListener() {
	fmt.Println("|||||||||||------ USER-PROFILE LISTENER STARTED! |||||||||||------")
	var (
		USER_PROFILE_STREAM = "user_profile_stream"
	)

	for {
		streams, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{USER_PROFILE_STREAM, "$"},
			Block:   0,
		}).Result()
		if err != nil {
			log.Printf("err reading from stream: %s || error: %s\n", USER_PROFILE_STREAM, err)
		}

		for _, stream := range streams {
			fmt.Println("stream: ", stream)
			for _, message := range stream.Messages {
				fmt.Println("message: ", message)
				s.uniqueId = message.Values["unique_id"].(string)
				getPostsAndTotalPostsCount(s)
				log.Printf("listened event from stream: %s || event: %s", USER_PROFILE_STREAM, message)
			}
		}
	}
}

func getPostsAndTotalPostsCount(s *Server) {
	posts := []Post{
		{
			PostId: 1,
			Url:    "image1.jpg",
		},
		{
			PostId: 2,
			Url:    "image2.jpg",
		},
		{
			PostId: 3,
			Url:    "image3.jpg",
		},
	}

	event := &PostEvent{
		UniqueId:   s.uniqueId,
		TotalPosts: 782,
		Posts:      posts,
	}
	s.PostsChan <- event
}
