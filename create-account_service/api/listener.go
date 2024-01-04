package api

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func (s *Server) UserProfileListener() {
	var (
		USER_PROFILE_STREAM = "user_profile_stream"
	)

	for {
		streams, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{USER_PROFILE_STREAM, "$"},
			Block:   0,
		}).Result()
		if err != nil {
			log.Printf("err reading from stream: %s || err: %s", USER_PROFILE_STREAM, err)
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				s.uniqueId = message.Values["unique_id"].(string)
				if err := getAcountInfo(s); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func getAcountInfo(s *Server) error {
	// parsedUUID, err := uuid.Parse(s.uniqueId)
	// if err != nil {
	// 	fmt.Errorf("err failed to parse unique_id: %s", err)
	// }

	event := &Account{
		Username: "Abdul Rehman",
		ImageUrl: "profile.jpg",
	}

	s.accountChan <- event
	return nil
}
