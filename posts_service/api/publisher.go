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

// //////////////////////////////////////////////////
// //////////////////////////////////////////////////
// //////////////////////////////////////////////////
// ///////-------PUBLISHER /////////////-------
func (s *Server) Publisher() error {
	fmt.Println("|||||||||||------ POSTS PUBLISHER STARTED! |||||||||||------")
	for {
		select {
		case event := <-s.postsChan:
			if err := publishPosts(s.redisClient, s.uniqueId, event); err != nil {
				fmt.Println("error publishing to stream: ", err)
			}
		case postsIds := <-s.postsIdsChan:
			if err := publishPostsIds(s.redisClient, postsIds); err != nil {
				fmt.Println("error publishing to stream: ", err)

			}
		}
	}
}

// helper function
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

// helper function
func publishPostsIds(redisClient *redis.Client, postsIds []int32) error {
	var (
		POST_IDS_STREAM = "posts:ids"
	)

	_, err := redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Stream: POST_IDS_STREAM,
		ID:     "*",
		Values: map[string]interface{}{
			"postsIds": postsIds,
		},
	}).Result()

	if err != nil {
		return fmt.Errorf("err adding event to the stream: %s", err)
	}
	return nil
}
