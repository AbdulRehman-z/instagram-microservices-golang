package main

import (
	"database/sql"
	"log/slog"

	"github.com/AbdulRehman-z/instagram-microservices/auth_service/cmd/api"
	db "github.com/AbdulRehman-z/instagram-microservices/auth_service/cmd/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/cmd/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		slog.Error("Error loading config: ", err)
	}

	conn, err := sql.Open(config.DB_DRIVER, config.DB_URL)
	if err != nil {
		slog.Error("Cannot connect to DB: ", err)
	}
	conn.Ping()
	store := db.NewStore(conn)

	server, err := api.NewServer(*config, store)
	if err != nil {
		slog.Error("Cannot create server: ", err)
	}

	err = server.Start(config.LISTEN_ADDR)
	if err != nil {
		slog.Error("Failed to start server: ", err)
	}
}

//
