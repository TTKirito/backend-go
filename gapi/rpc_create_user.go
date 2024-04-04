package gapi

import (
	"context"

	db "github.com/TTKirito/backend-go/db/sqlc"
	"github.com/TTKirito/backend-go/pb"
	"github.com/TTKirito/backend-go/utils"
	"github.com/TTKirito/backend-go/val"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	validations := validateCreateUserRequest(req)

	if validations != nil {
		return nil, invalidArgument(validations)
	}

	hashedPassword, err := utils.HashedPassword(req.GetPassword())

	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			switch pgErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "user already exists")
			}

		}

		return nil, status.Errorf(codes.Internal, "cannot create user")
	}

	rsp := &pb.CreateUserResponse{
		User: convertedUser(user),
	}

	return rsp, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	if err := val.ValidateFullName(req.GetFullName()); err != nil {
		violations = append(violations, fieldViolation("full_name", err))
	}
	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	return
}
