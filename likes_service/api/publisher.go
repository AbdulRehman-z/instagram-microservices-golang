package api

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func (s *Server) PostsLikesPublisher() {
	var (
		POST_LIKES_STREAM = "post_likes_stream"
	)

	for {
		likes := <-s.postlikesChan
		if err := s.redisClient.XAdd(context.Background(), &redis.XAddArgs{
			Stream: POST_LIKES_STREAM,
			Values: map[string]interface{}{
				"post_likes": likes,
			},
		}).Err(); err != nil {
			log.Println("Error while publishing post likes: ", err)
			continue
		}
	}
}

func (s *Server) CommentsLikesPublisher() {
	var (
		COMMENT_LIKES_STREAM = "comment_likes_stream"
	)

	for {
		likes := <-s.commentlikesChan
		if err := s.redisClient.XAdd(context.Background(), &redis.XAddArgs{
			Stream: COMMENT_LIKES_STREAM,
			Values: map[string]interface{}{
				"comment_likes": likes,
			},
		}).Err(); err != nil {
			log.Println("Error while publishing comment likes: ", err)
			continue
		}
	}
}
