package api

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func (s *Server) Publisher() {
	fmt.Println("|||||||||||---------- Starting |comments:count| publisher ----------|||||||||||")
	for {
		select {
		case commentsCount := <-s.commentsCountChan:
			if err := addCommentsCount(s, commentsCount); err != nil {
				log.Println("Error while publishing comment likes: ", err)
				continue
			}
		case ids := <-s.commentsIds:
			if err := addCommentsIds(s, ids); err != nil {
				log.Println("Error while publishing comment likes: ", err)
				continue
			}
		}
	}
}

func addCommentsCount(s *Server, commentsCount []CommentsCount) error {
	var (
		COMMENTS_COUNT_STREAM = "comments:count"
	)

	if err := s.redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Stream: COMMENTS_COUNT_STREAM,
		Values: map[string]interface{}{
			"comments_count": commentsCount,
		},
	}).Err(); err != nil {
		log.Println("Error while publishing comment likes: ", err)
	}

	return nil
}

func addCommentsIds(s *Server, ids []int32) error {
	var (
		COMMENTS_IDS_STREAM = "comments:ids"
	)

	if err := s.redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Stream: COMMENTS_IDS_STREAM,
		Values: map[string]interface{}{
			"comments_ids": ids,
		},
	}).Err(); err != nil {
		log.Println("Error while publishing comment likes: ", err)
	}

	return nil
}
