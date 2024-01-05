package api

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func (s *Server) PostsLikesPublisher() {
	var (
		LIKES_POSTS_STREAM = "likes:posts"
	)

	for {
		likes := <-s.postlikesChan
		if err := s.redisClient.XAdd(context.Background(), &redis.XAddArgs{
			Stream: LIKES_POSTS_STREAM,
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
		LIKES_COMMENTS_STREAM = "likes:comments"
	)

	for {
		likes := <-s.commentlikesChan
		if err := s.redisClient.XAdd(context.Background(), &redis.XAddArgs{
			Stream: LIKES_COMMENTS_STREAM,
			Values: map[string]interface{}{
				"comment_likes": likes,
			},
		}).Err(); err != nil {
			log.Println("Error while publishing comment likes: ", err)
			continue
		}
	}
}
