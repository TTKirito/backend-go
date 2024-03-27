package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomParticipant(t *testing.T) Participant {
	event := createRandomEvent(t)
	account := createRandomAccount(t)
	arg := CreateParticipantParams{
		Event:   event.ID,
		Account: account.ID,
	}

	participant, err := testQueries.CreateParticipant(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Event, participant.Event)
	require.Equal(t, arg.Account, participant.Account)

	return participant

}

func TestCreateParticipant(t *testing.T) {
	createRandomParticipant(t)
}

func TestDeleteParticipant(t *testing.T) {
	participant := createRandomParticipant(t)
	err := testQueries.DeleteParticipant(context.Background(), participant.ID)
	require.NoError(t, err)
}

func TestListParticipant(t *testing.T) {
	event := createRandomEvent(t)

	for i := 0; i < 10; i++ {
		account := createRandomAccount(t)
		arg := CreateParticipantParams{
			Event:   event.ID,
			Account: account.ID,
		}

		participant, err := testQueries.CreateParticipant(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, participant)
	}

	arg := ListParticipantParams{
		Limit:  5,
		Offset: 1,
		Event:  event.ID,
	}

	participants, err := testQueries.ListParticipant(context.Background(), arg)
	fmt.Println(participants)
	require.NoError(t, err)
	require.Len(t, participants, 5)

	for _, participant := range participants {
		require.NotEmpty(t, participant)
	}
}
