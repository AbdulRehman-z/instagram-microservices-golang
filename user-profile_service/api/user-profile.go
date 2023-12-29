package api

import (
	"context"
	"sync"

	"github.com/AbdulRehman-z/instagram-microservices/user-profile_service/token"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type Profile struct {
	Username       string `json:"username"`
	ImageUrl       string `json:"image_url"`
	TotlaFollowers int32  `json:"total_followers"`
	TotalFollowing int32  `json:"total_following"`
	TotalPosts     int32  `json:"total_posts"`
	Posts          []Post `json:"posts"`
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
	// send get followers count event to followers service using redis-streams
	// send get following count event to following service
	// send get Likes count event to likes service
	// send get posts count event to posts service
	// send get account event to account service
	payload := ctx.Locals(authorizationPayloadKey).(*token.Payload)

}

func fetchUserAccount(s *Server, wg *sync.WaitGroup, ch chan<- *Account, uniqueId string) error {
	defer wg.Done()

	// send get account event to account service
	res, err := s.redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Stream: "account",
		MaxLen: 0,
		ID:     "",
		Values: map[string]interface{}{
			"unique_id": uniqueId,
		},
	}).Result()
	if err != nil {
		return err
	}

	return nil
}
