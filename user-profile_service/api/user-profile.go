package api

import (
	"sync"

	"github.com/AbdulRehman-z/instagram-microservices/user-profile_service/token"
	"github.com/gofiber/fiber/v2"
)

type Profile struct {
	UserInfo                    Account `json:"user_info"`
	TotlaFollowersAndFollowings string  `json:"total_followers_and_followings"`
	TotalPosts                  int64   `json:"total_posts"`
	Posts                       []Post  `json:"posts"`
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
	followingsAndFollowersCount := make(chan string)
	postsCount := make(chan int64)
	posts := make(chan []Post)
	account := make(chan Account)

	s.Publisher(uniqueId.String())

	wg := sync.WaitGroup{}
	wg.Add(3)
	go GetFollowersAndFollowinCount(s, &wg, followingsAndFollowersCount)
	go GetPostsAndPostsCount(s, &wg, postsCount, posts)
	go GetAccount(s, &wg, account)
	wg.Wait()

	profile := Profile{
		UserInfo:                    <-account,
		TotlaFollowersAndFollowings: <-followingsAndFollowersCount,
		TotalPosts:                  <-postsCount,
		Posts:                       <-posts,
	}

	return ctx.Status(fiber.StatusOK).JSON(profile)
}

func GetFollowersAndFollowinCount(s *Server, wg *sync.WaitGroup, followingsAndFollowersCount chan string) {
	defer wg.Done()

loop:
	for {
		select {
		case count := <-s.totalFollowingsAndFollowersChan:
			followingsAndFollowersCount <- count
			break loop
		}
	}
}

func GetPostsAndPostsCount(s *Server, wg *sync.WaitGroup, postsCount chan int64, postsForUser chan []Post) {
	defer wg.Done()

	for {
		select {
		case count := <-s.totalPostsChan:
			postsCount <- count
		case posts := <-s.postsChan:
			postsForUser <- posts
		}
	}
}

func GetAccount(s *Server, wg *sync.WaitGroup, account chan Account) {
	defer wg.Done()

loop:
	for {
		select {
		case acc := <-s.accountChan:
			account <- acc
			break loop
		}
	}
}
