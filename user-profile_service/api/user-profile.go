package api

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/AbdulRehman-z/instagram-microservices/user-profile_service/token"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type Profile struct {
	UserInfo       Account `json:"user_info"`
	TotlaFollowers int32   `json:"total_followers"`
	TotalFollowing int32   `json:"total_following"`
	TotalPosts     int32   `json:"total_posts"`
	Posts          []Post  `json:"posts"`
}

type Post struct {
	PostId   int32  `json:"post_id"`
	ImageUrl string `json:"image_url"`
}

type Account struct {
	Username string `json:"username"`
	ImageUrl string `json:"image_url"`
}

func (s *Server) UserProfile(ctx *fiber.Ctx) error {
	payload := ctx.Locals(authorizationPayloadKey).(*token.Payload)
	uniqueId := payload.UniqueId

	totalFollowersChan := make(chan int32)
	totalFollowingChan := make(chan int32)
	totalPostsChan := make(chan int32)
	postsChan := make(chan []Post)
	accountChan := make(chan Account)

	wg := sync.WaitGroup{}
	go GetFollowersAndFollowinCount(s, uniqueId.String(), &wg, totalFollowersChan, totalFollowingChan)
	go GetPostsAndPostsCount(s, uniqueId.String(), &wg, totalPostsChan, postsChan)
	go GetAccount(s, uniqueId.String(), &wg, accountChan)
	wg.Add(5)
	wg.Wait()

	profile := Profile{
		UserInfo:       <-accountChan,
		TotlaFollowers: <-totalFollowersChan,
		TotalFollowing: <-totalFollowingChan,
		TotalPosts:     <-totalPostsChan,
		Posts:          <-postsChan,
	}

	return ctx.Status(fiber.StatusOK).JSON(profile)
}

func GetFollowersAndFollowinCount(s *Server, uniqueId string, wg *sync.WaitGroup, totalFollowersChan chan int32, totalFollowingChan chan int32) {
	defer wg.Done()
	var (
		FOLLOWERS_FOLLOWING_COUNT = "followers_following_count"
	)

	if err := s.publish(uniqueId); err != nil {
		log.Printf("error publishing to stream: %d", err)
	}

	// wait for the response from the consumer
	for {
		response, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{FOLLOWERS_FOLLOWING_COUNT, "0"}, //
			Block:   0,
		}).Result()
		if err != nil {
			log.Printf("error reading from stream: %d", err)
			continue
		}

		for _, message := range response[0].Messages {
			if message.Values["unique_id"] == uniqueId {
				totalFollowersChan <- message.Values["total_followers"].(int32)
				totalFollowingChan <- message.Values["total_following"].(int32)
				fmt.Printf("total followers: %d", message.Values["total_followers"].(int32))
				return
			}
		}
		break
	}
}

func GetPostsAndPostsCount(s *Server, uniqueId string, wg *sync.WaitGroup, totalPostsChan chan int32, posts chan []Post) {
	defer wg.Done()
	var (
		POSTS = "posts"
	)

	if err := s.publish(uniqueId); err != nil {
		log.Printf("error publishing to stream: %d", err)
	}

	// wait for the response from the consumer
	for {
		response, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{POSTS, "0"},
			Block:   0,
		}).Result()
		if err != nil {
			log.Printf("error reading from stream: %d", err)
			continue
		}

		for _, message := range response[0].Messages {
			if message.Values["unique_id"] == uniqueId {
				totalPostsChan <- message.Values["total_posts"].(int32)
				posts <- message.Values["posts"].([]Post)
				fmt.Printf("total posts: %d", message.Values["total_posts"].(int32))
				return
			}
		}
	}
}

func GetAccount(s *Server, uniqueId string, wg *sync.WaitGroup, accountChan chan Account) {
	defer wg.Done()
	var (
		ACCOUNT = "account"
	)

	if err := s.publish(uniqueId); err != nil {
		log.Printf("error publishing to stream: %d", err)
	}

	// wait for the response from the consumer
	for {
		response, err := s.redisClient.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{ACCOUNT, "0"}, //
			Block:   0,
		}).Result()
		if err != nil {
			log.Printf("error reading from stream: %d", err)
			continue
		}

		for _, message := range response[0].Messages {
			if message.Values["unique_id"] == uniqueId {
				accountChan <- Account{
					Username: message.Values["username"].(string),
					ImageUrl: message.Values["image_url"].(string),
				}
				fmt.Printf("username: %s", message.Values["username"].(string))
				return
			}
		}
	}
}
