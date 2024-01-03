package api

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (s *Server) Listener() error {
	fmt.Println("||||||||||||| LISTENER ||||||||||||||")
	var (
		USER_PROFILE_STREAM = "user_profile_stream"
		// USER_PROFILE_GROUP      = "user_profile_group"
		// USER_PROFILE_CONSUMER_1 = "user_profile_consumer_1"
	)

	// _, err := s.redisClient.XGroupCreate(context.Background(), USER_PROFILE_STREAM, USER_PROFILE_GROUP, "0").Result()
	// if err != nil {
	// 	if err.Error() == "BUSYGROUP Consumer Group name already exists" {
	// 		fmt.Printf("group already exists: %s\n", USER_PROFILE_GROUP)
	// 	} else {
	// 		return fmt.Errorf("err creating group: %d", err)
	// 	}
	// }

	for {
		streams, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{USER_PROFILE_STREAM, "$"},
			Block:   0,
		}).Result()
		if err != nil {
			return fmt.Errorf("error reading from stream: %d", err)
		}

		if len(streams) > 0 {
			for _, stream := range streams {
				for _, message := range stream.Messages {
					s.uniqueId = message.Values["unique_id"].(string)
					getFollowersAndFollowingCountFromStore(s, s.uniqueId)
					log.Printf("listened event from stream: %s || event: %s", USER_PROFILE_STREAM, message)
				}
			}
		}
	}
}

func getFollowersAndFollowingCountFromStore(s *Server, uniqueId string) {
	uniqueIdUUID, err := uuid.Parse(s.uniqueId)
	if err != nil {
		log.Println("err parsing uuid: ", err)
	}
	count, err := s.store.GetFollowingAndFollowersCount(context.Background(), uniqueIdUUID)
	if err != nil {
		log.Println("err getting followers and following count: ", err)
	}
	s.followersFollowingCountChan <- fmt.Sprintf("%d", count)
}
