package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/TTKirito/backend-go/api"
	db "github.com/TTKirito/backend-go/db/sqlc"
	"github.com/TTKirito/backend-go/gapi"
	"github.com/TTKirito/backend-go/pb"
	"github.com/TTKirito/backend-go/utils"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	runDBMigration(config.MIGRATION_PATH, config.DBSOURCE)
	store := db.NewStore(conn)
	go runGinServer(config, store)
	runGRPCServer(config, store)
}

func runDBMigration(migrationPath string, dbSource string) {
	migration, err := migrate.New(migrationPath, dbSource)

	if err != nil {
		log.Fatal("cannot create migration instance:", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("cannot to run migration up", err)
	}

	log.Printf("db migrate succsessfully")
}

func runGRPCServer(config utils.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEventServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)

	if err != nil {
		log.Fatal("cannot listen server:", err)
	}

	log.Printf("start GRPC server at %s", listener.Addr())
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}

func runGinServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	err = server.Start((config.HTTPServerAddress))

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

// export PATH="$PATH:$(go env GOPATH)/bin"
