package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/TTKirito/backend-go/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Position: Positions(utils.RandomPosition()),
		Gender:   Genders(utils.RandomGender()),
		Dob:      utils.RandomDob(),
		Status:   Status(utils.RandomStatus()),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Position, account.Position)
	require.Equal(t, arg.Gender, account.Gender)
	require.WithinDuration(t, arg.Dob, account.Dob, time.Second)
	require.Equal(t, arg.Status, account.Status)

	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.ID)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)

	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, account.Position, account2.Position)
	require.Equal(t, account.Gender, account2.Gender)
	require.WithinDuration(t, account.Dob, account2.Dob, time.Second)
	require.Equal(t, account.Status, account2.Status)

	require.NotZero(t, account2.CreatedAt)
	require.NotZero(t, account2.ID)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateAccountParams{
		Position: Positions(utils.RandomPosition()),
		Gender:   Genders(utils.RandomGender()),
		Dob:      utils.RandomDob(),
		Status:   Status(utils.RandomStatus()),
		ID:       account.ID,
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, arg.Position, account2.Position)
	require.Equal(t, arg.Gender, account2.Gender)
	require.WithinDuration(t, arg.Dob, account2.Dob, time.Second)
	require.Equal(t, arg.Status, account2.Status)

	require.NotZero(t, account2.CreatedAt)
	require.NotZero(t, account2.ID)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)

	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}
