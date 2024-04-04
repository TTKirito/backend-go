package gapi

import (
	"fmt"

	db "github.com/TTKirito/backend-go/db/sqlc"
	"github.com/TTKirito/backend-go/pb"
	"github.com/TTKirito/backend-go/token"
	"github.com/TTKirito/backend-go/utils"
)

type Server struct {
	pb.UnimplementedEventServer
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token marker %v", err)
	}

	return &Server{config: config, store: store, tokenMaker: tokenMaker}, nil
}
