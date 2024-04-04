package main

import (
	"database/sql"
	"log"

	"github.com/TTKirito/backend-go/api"
	db "github.com/TTKirito/backend-go/db/sqlc"
	"github.com/TTKirito/backend-go/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDRIVER, config.DBSOURCE)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	err = server.Start((config.ServerAddress))

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

// export PATH="$PATH:$(go env GOPATH)/bin"
