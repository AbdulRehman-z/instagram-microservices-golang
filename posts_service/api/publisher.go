package api

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Post struct {
	PostId int64
	Url    string
}

type PostEvent struct {
	UniqueId   string
	TotalPosts int64
	Posts      []Post
}

func (pe PostEvent) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer

	err := gob.NewEncoder(&buf).Encode(pe.Posts)
	if err != nil {
		return nil, fmt.Errorf("err encoding: %d", err)
	}
	return buf.Bytes(), nil
}

func (s *Server) Publisher() error {
	fmt.Println("|||||||||||------ POSTS PUBLISHER STARTED! |||||||||||------")
	for posts := range s.PostsChan {
		if err := publishPosts(s.redisClient, s.uniqueId, posts); err != nil {
			fmt.Println("error publishing to stream: ", err)
		}
		fmt.Println("----------------------------")
		fmt.Println("posts published to posts_stream: ", posts)
		fmt.Println("----------------------------")
	}
	return nil
}

func publishPosts(redisClient *redis.Client, uniqueId string, event *PostEvent) error {
	var (
		POSTS_STREAM = "posts_stream"
	)

	binary, err := event.MarshalBinary()
	if err != nil {
		return fmt.Errorf("err marshalling event: %s", err)
	}

	_, err = redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Stream: POSTS_STREAM,
		ID:     "*",
		Values: map[string]any{
			"TotalPosts": event.TotalPosts,
			"Posts":      binary,
		},
	}).Result()

	if err != nil {
		return fmt.Errorf("err adding event to the stream: %s", err)
	}
	return nil
}
