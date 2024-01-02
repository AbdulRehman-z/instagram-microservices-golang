package api

import (
	"context"
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (s *Server) Listener() error {
	fmt.Println("||||||||||||| LISTENER ||||||||||||||")
	wg := sync.WaitGroup{}
	var (
		USER_PROFILE_STREAM     = "user_profile_stream"
		USER_PROFILE_GROUP      = "user_profile_group"
		USER_PROFILE_CONSUMER_1 = "user_profile_consumer_1"
	)

	_, err := s.redisClient.XGroupCreate(context.Background(), USER_PROFILE_STREAM, USER_PROFILE_GROUP, "0").Result()
	if err != nil {
		if err.Error() == "BUSYGROUP Consumer Group name already exists" {
			fmt.Printf("group already exists: %s\n", USER_PROFILE_GROUP)
		} else {
			return fmt.Errorf("err creating group: %d", err)
		}
	}

	for {
		streams, err := s.redisClient.XReadGroup(context.Background(), &redis.XReadGroupArgs{
			Streams:  []string{USER_PROFILE_STREAM, ">"},
			Group:    USER_PROFILE_GROUP,
			Consumer: USER_PROFILE_CONSUMER_1,
			Block:    0,
			NoAck:    false,
		}).Result()
		if err != nil {
			return fmt.Errorf("error reading from stream: %d", err)
		}

		if len(streams) > 0 {
			for _, stream := range streams {
				for _, message := range stream.Messages {
					s.uniqueId = message.Values["unique_id"].(string)
					s.redisClient.XAck(context.Background(), USER_PROFILE_STREAM, stream.Messages[0].ID)
					go getFollowersCountFromStore(s, &wg, s.uniqueId)
					// go getFollowingCountFromStore(s, &wg, s.followingCountChan, s.uniqueId)
					wg.Add(1)
				}
			}
		}
	}
}

func getFollowersCountFromStore(s *Server, wg *sync.WaitGroup, uniqueId string) {
	defer func() {
		wg.Done()
	}()
	uniqueIdUUID, err := uuid.Parse(s.uniqueId)
	if err != nil {
		log.Errorf("err parsing uuid: %d", uniqueIdUUID)
	}
	count, err := s.store.GetFollowingAndFollowersCount(context.Background(), uniqueIdUUID)
	if err != nil {
		log.Errorf("err getting following count from db: %d", err)
	}
	s.followersFollowingCountChan <- fmt.Sprintf("%d", count)
}
