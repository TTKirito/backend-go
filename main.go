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

	store := db.NewStore(conn)
	runGRPCServer(config, store)
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
