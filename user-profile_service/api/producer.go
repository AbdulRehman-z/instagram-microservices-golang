package api

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func (s *Server) publish(uniqueId string) error {
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
	if err != nil {
		return fmt.Errorf("error adding to stream:")
	}
	if err != nil {
		return fmt.Errorf("error adding to stream: %d", err)
	}
	log.Printf("event added to stream! stream_name: %d ----- response: %d", USER_PROFILE_STREAM, res)

	return nil
}
