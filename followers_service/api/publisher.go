package api

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func (s *Server) Publisher() error {
	fmt.Println("||||||||||||| PUBLISHER ||||||||||||||")
	for {
		select {
		case followersCount := <-s.followersCountChan:
			fmt.Println("----------------------------")
			fmt.Println("followers count: ", followersCount)
			fmt.Println("----------------------------")
			if err := s.publishFollowersCount(s.uniqueId, followersCount); err != nil {
				log.Printf("error publishing to stream: %d", err)
			}
		case followingCount := <-s.followingCountChan:
			fmt.Println("----------------------------")
			fmt.Println("following count: ", followingCount)
			fmt.Println("----------------------------")
			if err := s.publishFollowingCount(s.uniqueId, followingCount); err != nil {
				log.Printf("error publishing to stream: %d", err)
			}
		}
	}
}

func (s *Server) publishFollowersCount(uniqueId string, followersCount int64) error {
	var (
		FOLLOWERS_FOLLOWING_COUNT = "followers_following_count"
	)
	res, err := s.redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Values: map[string]interface{}{
			"unique_id":       uniqueId,
			"total_followers": followersCount,
		},
		ID:     "*",
		Stream: FOLLOWERS_FOLLOWING_COUNT,
	}).Result()
	if err != nil {
		return fmt.Errorf("error adding to stream: %d", err)
	}
	log.Printf("added to stream: %s", res)
	return nil
}

func (s *Server) publishFollowingCount(uniqueId string, followingCount int64) error {
	var (
		FOLLOWERS_FOLLOWING_COUNT = "followers_following_count"
	)
	res, err := s.redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Values: map[string]interface{}{
			"unique_id":       uniqueId,
			"total_following": followingCount,
		},
		ID:     "*",
		Stream: FOLLOWERS_FOLLOWING_COUNT,
	}).Result()
	if err != nil {
		return fmt.Errorf("error adding to stream: %d", err)
	}
	log.Printf("added to stream: %s", res)
	return nil
}
