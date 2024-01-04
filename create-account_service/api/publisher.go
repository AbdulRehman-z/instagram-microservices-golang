package api

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type Account struct {
	Username string
	ImageUrl string
}

func (a *Account) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer

	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(a.Username)
	if err != nil {
		return nil, fmt.Errorf("err encoding: %d", err)
	}
	err = encoder.Encode(a.ImageUrl)
	if err != nil {
		return nil, fmt.Errorf("err encoding: %d", err)
	}

	fmt.Println(buf.Bytes())
	return buf.Bytes(), nil
}

func (s *Server) Publisher() {
	for account := range s.accountChan {
		if err := publishAccount(s, account); err != nil {
			log.Printf("error publishing event: %d\n", err)
		}
		log.Printf("event added to stream: %s || event: %s\n", "account_stream", account)
	}
}

func publishAccount(s *Server, event *Account) error {
	var (
		ACCOUNT_STREAM = "account_stream"
	)
	binary, err := event.MarshalBinary()
	if err != nil {
		return err
	}
	_, err = s.redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Stream: ACCOUNT_STREAM,
		ID:     "*",
		Values: map[string]any{
			"unique_id": s.uniqueId,
			"account":   binary,
		},
	}).Result()
	if err != nil {
		return fmt.Errorf("err adding to stream: %d", err)
	}
	return nil
}
