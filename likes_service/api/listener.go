package api

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type CommentLikes struct {
	Comment_id int32
	Likes      int64
}

type PostLikes struct {
	Post_id int32
	likes   int64
}

// /////////////////////////////////////////////////////
// /////////////////////////////////////////////////////
// /////////////////////////////////////////////////////
// ///////////------- LISTENER /////////////-------
func (s *Server) PostsLikesListener() {
	fmt.Println("|||||||||||---------- Starting PostsLikesListener ----------|||||||||||")
	var (
		POST_IDS_STREAM = "posts:ids"
	)

	for {
		streams, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{POST_IDS_STREAM, "$"},
			Block:   0,
		}).Result()
		if err != nil {
			log.Println("Error while reading stream: ", err)
			continue
		}

		for _, message := range streams {
			for _, msg := range message.Messages {
				postIds := msg.Values["posts_ids"].([]int32)
				if err := getPostLikes(s, postIds); err != nil {
					log.Println(err)
					continue
				}
			}
		}
	}
}

func getPostLikes(s *Server, Ids []int32) error {
	array := make([]PostLikes, 0)

	for _, id := range Ids {
		res, err := s.store.GetPostLikesCount(context.Background(), id)
		if err != nil {
			return fmt.Errorf("error while getting post likes count: %s", err)
		}

		postLike := PostLikes{
			Post_id: res.PostID,
			likes:   res.Count,
		}

		// append each comment likes to an array and then send that array to the channel
		array = append(array, postLike)
	}

	s.postlikesChan <- array
	return nil
}

// /////////////////////////////////////////////////////
// /////////////////////////////////////////////////////
// /////////////////////////////////////////////////////
// ///////////------- LISTENER /////////////-------
func (s *Server) CommentsLikesListener() {
	fmt.Println("|||||||||||---------- Starting CommentsLikesListener ----------|||||||||||")
	var (
		COMMENTS_LIKES_STREAM = "comments:ids"
	)

	for {
		streams, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{COMMENTS_LIKES_STREAM, "$"},
			Block:   0,
		}).Result()
		if err != nil {
			log.Println("Error while reading stream: ", err)
			continue
		}
		// get comments Ids from array
		for _, stream := range streams {
			for _, message := range stream.Messages {
				commentIds := message.Values["comments_ids"].([]int32)
				getCommentsLikes(s, commentIds)
			}
		}
	}
}

func getCommentsLikes(s *Server, Ids []int32) error {
	array := make([]CommentLikes, 0)
	for _, id := range Ids {
		res, err := s.store.GetCommentLikesCount(context.Background(), id)
		if err != nil {
			return fmt.Errorf("error while getting comment likes count: %s", err)
		}

		commentLikes := CommentLikes{
			Comment_id: res.CommentID,
			Likes:      res.Count,
		}

		// append each comment likes to an array and then send that array to the channel
		array = append(array, commentLikes)
	}

	s.commentlikesChan <- array
	return nil
}
