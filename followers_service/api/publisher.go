package api

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func (s *Server) Publisher() error {
	fmt.Println("||||||||||||| PUBLISHER ||||||||||||||")
	for followersFollowingCount := range s.followersFollowingCountChan {
		fmt.Println("----------------------------")
		fmt.Println("followers and following count published: ", followersFollowingCount)
		fmt.Println("----------------------------")
		if err := s.publishFollowersCount(s.uniqueId, followersFollowingCount); err != nil {
			log.Printf("error publishing to stream: %d", err)
		}
	}
	return nil
}

func (s *Server) publishFollowersCount(uniqueId string, followersFollowingCount string) error {
	var (
		FOLLOWERS_FOLLOWING_COUNT_STREAM = "followers_following_count_stream"
	)
	_, err := s.redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Values: map[string]interface{}{
			"unique_id":                 uniqueId,
			"total_followers_following": followersFollowingCount,
		},
		ID:     "*",
		Stream: FOLLOWERS_FOLLOWING_COUNT_STREAM,
	}).Result()
	if err != nil {
		return fmt.Errorf("error adding to stream: %d", err)
	}
	return nil
}
