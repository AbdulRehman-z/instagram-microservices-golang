package cache

import "github.com/redis/go-redis/v9"

// type RedisClient struct {
// 	Host     string
// 	Password string
// 	DB       string
// }

func NewRedisClient(Host, Port, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     Host + ":" + Port,
		Password: password,
		DB:       db,
	})

	return client
}
