package db

import (
	"context"
	"testing"
	"time"

	"github.com/TTKirito/backend-go/token"
	"github.com/TTKirito/backend-go/utils"
	"github.com/stretchr/testify/require"
)

func RandomSession(t *testing.T) Session {
	maker, err := token.NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)
	user := createRandomUser(t)
	token, payload, err := maker.CreateToken(user.Username, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	payload, err = maker.VerifyToken(token)
	require.NotEmpty(t, payload)
	require.NoError(t, err)

	refreshToken, payload, err := maker.CreateToken(user.Username, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	arg := CreateSessionParams{
		ID:           payload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    "",
		IsBlocked:    false,
		ClientIp:     "",
		ExpiredAt:    payload.ExpiredAt,
	}

	session, err := testQueries.CreateSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session)
	return session
}

func TestCreateSession(t *testing.T) {
	RandomSession(t)
}
