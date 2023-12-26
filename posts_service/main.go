package main

import (
	"database/sql"
	"log"

	"github.com/AbdulRehman-z/instagram-microservices/posts_service/api"
	db "github.com/AbdulRehman-z/instagram-microservices/posts_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/posts_service/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Printf("Error loading config: %v", err)
	}

	conn, err := sql.Open(config.DB_DRIVER, config.DB_URL)
	if err != nil {
		log.Printf("Cannot connect to DB: %v", err)
	}
	util.RunMigration(config.DB_MIGRATION_URL, config.DB_URL)
	if err := conn.Ping(); err != nil {
		log.Printf("Cannot ping DB: %v", err)
	}

	store := db.NewStore(conn)
	run(config, store)
}

func run(config *util.Config, store db.Store) {
	server, err := api.NewServer(*config, store)
	if err != nil {
		log.Fatalf("failed to initiate server: %d", err)
	}
	err = server.Start(config.LISTEN_ADDR)
	if err != nil {
		log.Fatalf("failed to start server: %d", err)
	}
}
