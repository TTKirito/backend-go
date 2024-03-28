package db

import (
	"context"
	"testing"

	"github.com/TTKirito/backend-go/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashedPassword(utils.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       utils.RandomOwner(),
		FullName:       utils.RandomOwner(),
		HashedPassword: hashedPassword,
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateuser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	user1, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.Equal(t, user.Username, user1.Username)
	require.Equal(t, user.FullName, user1.FullName)
	require.Equal(t, user.HashedPassword, user1.HashedPassword)
	require.Equal(t, user.Email, user1.Email)
	require.NotZero(t, user1.CreatedAt)
}
