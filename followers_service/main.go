package main

import (
	"database/sql"
	"log"

	"github.com/AbdulRehman-z/instagram-microservices/followers_service/api"
	db "github.com/AbdulRehman-z/instagram-microservices/followers_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/followers_service/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := sql.Open(config.DB_DRIVER, config.DB_URL)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store := db.NewStore(conn)
	run(config, store)
}

func run(config *util.Config, store db.Store) error {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}
	go server.Listener()
	go server.Publisher()
	err = server.Start()
	if err != nil {
		return err
	}
	return nil
}
