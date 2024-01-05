package api

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type CommentsCount struct {
	POST_ID int32
	Likes   int64
}

type CommentsLikes struct {
	Comment_id int32
	Likes      int64
}

// ///////////////////////////////////////////////////////////
// ///////////////////////////////////////////////////////////
// ///////////////////////////////////////////////////////////
// /////////------- LISTENER /////////////-------
func (s *Server) PostsIdsListener() {
	fmt.Println("|||||||||||---------- Starting |likes:comments| listener ----------|||||||||||")
	var (
		POST_IDS_STREAM = "posts:ids"
	)

	for {
		streams, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{POST_IDS_STREAM, "$"},
			Block:   0,
		}).Result()
		if err != nil {
			slog.Error("err reading from stream", "stream", POST_IDS_STREAM)
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				postIds := message.Values["posts_ids"].([]int32)
				if err := getCommentsCountForPost(s, postIds); err != nil {
					slog.Error("err getting comments count for post", "err", err)
				}
			}
		}
	}
}

func getCommentsCountForPost(s *Server, postIds []int32) error {
	array := make([]CommentsCount, 0)

	for _, id := range postIds {
		res, err := s.store.GetCommentsCount(context.Background(), id)
		if err != nil {
			return fmt.Errorf("err getting comments count: %s", err)
		}
		array = append(array, CommentsCount{POST_ID: id, Likes: res})
	}

	s.commentsCountChan <- array

	return nil
}

// //////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////
// ///////------- LISTENER /////////////-------
func (s *Server) CommentsLikesListener() {
	fmt.Println("|||||||||||---------- Starting |likes:comments| listener ----------|||||||||||")
	var (
		COMMENTS_LIKES_STREAM = "likes:comments"
	)

	for {
		streams, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{COMMENTS_LIKES_STREAM, "$"},
			Block:   0,
		}).Result()
		if err != nil {
			slog.Error("err reading from stream", "stream", COMMENTS_LIKES_STREAM)
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				commentsLikes := message.Values["comment_likes"].([]CommentsLikes)
				slog.Info("listened event from stream", "stream", COMMENTS_LIKES_STREAM, "event", commentsLikes)
			}
		}
	}
}
