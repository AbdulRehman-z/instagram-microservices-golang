package api

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

// /////////////////////////////////////////////////
// /////////////////////////////////////////////////
// /////////////////////////////////////////////////
// /////////------- USER-PROFILE LISTENER /////////////-------
func (s *Server) UserProfileListener() {
	fmt.Println("|||||||||||------ USER-PROFILE LISTENER STARTED! |||||||||||------")
	var (
		USER_PROFILE_STREAM = "user_profile_stream"
	)

	for {
		streams, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{USER_PROFILE_STREAM, "$"},
			Block:   0,
		}).Result()
		if err != nil {
			log.Printf("err reading from stream: %s || error: %s\n", USER_PROFILE_STREAM, err)
		}

		for _, stream := range streams {
			fmt.Println("stream: ", stream)
			for _, message := range stream.Messages {
				fmt.Println("message: ", message)
				s.uniqueId = message.Values["unique_id"].(string)
				getPostsAndTotalPostsCount(s)
				log.Printf("listened event from stream: %s || event: %s", USER_PROFILE_STREAM, message)
			}
		}
	}
}

func getPostsAndTotalPostsCount(s *Server) {
	posts := []Post{
		{
			PostId: 1,
			Url:    "image1.jpg",
		},
		{
			PostId: 2,
			Url:    "image2.jpg",
		},
		{
			PostId: 3,
			Url:    "image3.jpg",
		},
	}

	event := &PostEvent{
		UniqueId:   s.uniqueId,
		TotalPosts: 782,
		Posts:      posts,
	}
	s.postsChan <- event
}

// //////////////////////////////////////////////////////
// //////////////////////////////////////////////////////
// //////////////////////////////////////////////////////
// /////////------- POSTS-COMMENTS-COUNT LISTENER /////////////-------
func (s *Server) PostsCommentsCountListener() {
	fmt.Println("|||||||||||------ POSTS-COMMENTS-COUNT LISTENER STARTED! |||||||||||------")
	var (
		COMMENTS_COUNT_STREAM = "comments:count"
	)

	type CommentsCount struct {
		POST_ID  int32
		Comments int64
	}

	for {
		streams, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{COMMENTS_COUNT_STREAM, "$"},
			Block:   0,
		}).Result()
		if err != nil {
			log.Printf("err reading from stream: %s || error: %s\n", COMMENTS_COUNT_STREAM, err)
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				commentsCount := message.Values["comments_count"].([]CommentsCount)
				log.Printf("listened event from stream: %s || event: %d\n", COMMENTS_COUNT_STREAM, commentsCount)
			}
		}
	}
}

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
///////////------- POSTS-LIKES-COUNT LISTENER /////////////-------

func (s *Server) PostsLikesListener() {
	fmt.Println("|||||||||||------ POSTS-LIKES LISTENER STARTED! |||||||||||------")
	var (
		POST_LIKES_STREAM = "posts:likes"
	)

	type PostLikes struct {
		Post_id int32
		Likes   int64
	}

	for {
		streams, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{POST_LIKES_STREAM, "$"},
			Block:   0,
		}).Result()
		if err != nil {
			log.Printf("err reading from stream: %s || error: %s\n", POST_LIKES_STREAM, err)
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				postLikes := message.Values["post_likes"].([]PostLikes)
				log.Printf("listened event from stream: %s || event: %d\n", POST_LIKES_STREAM, postLikes)
			}
		}
	}
}
