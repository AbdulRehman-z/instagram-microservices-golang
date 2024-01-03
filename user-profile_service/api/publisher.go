package api

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func (s *Server) Publisher(uniqueId string) error {
	fmt.Println("|||||||||||||------- USER-PROFILE-LISTENER PUBLISHER STARTED -------||||||||||||||")
	var (
		USER_PROFILE_STREAM = "user_profile_stream"
	)
	res, err := s.redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Values: map[string]any{
			"unique_id": uniqueId,
		},
		ID:     "*",
		Stream: USER_PROFILE_STREAM,
	}).Result()
	fmt.Println(res)
	if err != nil {
		return fmt.Errorf("error adding to stream: %d", err)
	}
	log.Printf("added to stream: %s", res)
	return nil
}
