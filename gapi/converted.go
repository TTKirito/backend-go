package gapi

import (
	db "github.com/TTKirito/backend-go/db/sqlc"
	"github.com/TTKirito/backend-go/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertedUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
