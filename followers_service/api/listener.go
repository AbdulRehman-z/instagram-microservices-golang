package api

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (s *Server) Listener() error {
	wg := sync.WaitGroup{}
	var (
		USER_PROFILE_STREAM = "user_profile_stream"
	)

	res, err := s.redisClient.XGroupCreate(context.Background(), USER_PROFILE_STREAM, "user_profile_group", "0").Result()
	if err != nil {
		fmt.Printf("error creating group: %d", err)
	}
	fmt.Println("res: ", res)
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("||||||||||||| LISTENER ||||||||||||||")

		streams, err := s.redisClient.XReadGroup(context.Background(), &redis.XReadGroupArgs{
			Streams:  []string{USER_PROFILE_STREAM, ">"},
			Group:    "user_profile_group",
			Consumer: "followers_consumer",
			Block:    0,
			NoAck:    false,
		}).Result()

		if err != nil {
			return fmt.Errorf("error reading from stream: %d", err)
		}

		// // res, err := s.redisClient.XGroupCreate(context.Background(), USER_PROFILE_STREAM, "followers_group", "$").Result()
		// if err != nil {
		// 	return fmt.Errorf("error creating group: %d", err)
		// }

		// fmt.Println("res: ", res)

		// if len(streams) == 0 {
		// 	continue
		// }

		if len(streams) > 0 {
			for _, stream := range streams {
				fmt.Println(stream)
				for _, message := range stream.Messages {
					s.uniqueId = message.Values["unique_id"].(string)
					s.redisClient.XAck(context.Background(), USER_PROFILE_STREAM, stream.Messages[0].ID)
					go getFollowersCountFromStore(s, &wg, s.followersCountChan, s.uniqueId)
					// go getFollowingCountFromStore(s, &wg, s.followingCountChan, s.uniqueId)
					wg.Add(1)
				}
			}
		}
	}
}

func getFollowersCountFromStore(s *Server, wg *sync.WaitGroup, followersCountChan chan int64, uniqueId string) {
	fmt.Printf("Context: getFollowersCountFromStore, UniqueId: %s\n", uniqueId)

	defer func() {
		fmt.Println("terminating getFollowersCountFromStore")
		wg.Done()
	}()

	// uniqueIdUUID, err := uuid.Parse(s.uniqueId)
	// if err != nil {
	// 	log.Errorf("err parsing uuid: %d", err)
	// }
	// followersCount, err := s.store.GetFollowersCount(context.Background(), uniqueIdUUID)
	// if err != nil {
	// 	log.Errorf("err getting following count from db: %d", err)
	// }
	s.followersCountChan <- 20
	// close(s.followersCountChan)
}

func getFollowingCountFromStore(s *Server, wg *sync.WaitGroup, followingCountChan chan int64, uniqueId string) {
	fmt.Printf("Context: getFollowingCountFromStore, UniqueId: %s\n", uniqueId)
	defer func() {
		fmt.Println("terminating getFollowingCountFromStore")
		wg.Done()
	}()
	uniqueIdUUID, err := uuid.Parse(s.uniqueId)
	if err != nil {
		log.Errorf("err parsing uuid: %d", err)
	}
	followingCount, err := s.store.GetFollowingCount(context.Background(), uniqueIdUUID)
	if err != nil {
		log.Errorf("err getting following count from db: %d", err)
	}
	followingCountChan <- followingCount
}
