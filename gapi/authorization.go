package gapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/TTKirito/backend-go/token"
	"google.golang.org/grpc/metadata"
)

const (
	authorizationHeader     = "authorization"
	authorizationTypeBearer = "bearer"
)

func (server *Server) authorizeUser(ctx context.Context) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	values := md.Get(authorizationHeader)

	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	authorHeader := values[0]

	fields := strings.Fields(authorHeader)

	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization")
	}

	authorType := strings.ToLower(fields[0])
	if authorType != authorizationTypeBearer {
		return nil, fmt.Errorf("unsupported authorization type %s", authorType)
	}

	accessToken := fields[1]
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)

	}

	return payload, nil
}
