package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/TTKirito/backend-go/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateEventTx(t *testing.T) {
	store := NewStore(testDB)

	account := createRandomAccount(t)
	account2 := createRandomAccount(t)
	account3 := createRandomAccount(t)

	arg := CreateEventTxParams{
		Title:      utils.RandomString(12),
		StartTime:  utils.RandomTime(),
		EndTime:    utils.RandomTime().Add(30 * time.Minute),
		IsEmegency: utils.RandomEmegency(),
		Owner:      account.ID,
		Note:       utils.RandomString(12),
		Type:       EventTypes(utils.RandomEventType()),
		VisitType:  VisitTypes(utils.RandomVisitType()),
		Meeting:    "http://",
		Location: Location{
			Lat:    utils.RandomLatLong().Lat,
			Long:   utils.RandomLatLong().Long,
			Street: utils.RandomLatLong().Street,
		},
		Participants: []Participant{
			{ID: account2.ID}, {ID: account3.ID},
		},
	}

	n := 5
	errs := make(chan error)
	results := make(chan CreateEventTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CreateEventTx(context.Background(), arg)

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)

		event := result.Event
		require.Equal(t, event.Title, sql.NullString{String: arg.Title, Valid: true})
		require.WithinDuration(t, event.StartTime, arg.StartTime, time.Second)
		require.WithinDuration(t, event.EndTime, arg.EndTime, time.Second)
		require.Equal(t, event.IsEmegency, arg.IsEmegency)
		require.Equal(t, event.Owner, arg.Owner)
		require.Equal(t, event.Note, sql.NullString{String: arg.Note, Valid: true})
		require.Equal(t, event.Type, arg.Type)
		require.Equal(t, event.VisitType, arg.VisitType)
		require.Equal(t, event.Meeting, sql.NullString{String: arg.Meeting, Valid: true})

		require.NotZero(t, event.ID)
		require.NotZero(t, event.CreatedAt)

		location := result.Location
		require.Equal(t, location.Lat, arg.Location.Lat)
		require.Equal(t, location.Long, arg.Location.Long)
		require.Equal(t, location.Street, arg.Location.Street)
		require.Equal(t, location.Event, event.ID)
		require.NotZero(t, location.ID)
		require.NotZero(t, location.CreatedAt)

		participants := result.Participants
		for _, participant := range participants {
			require.Equal(t, participant.Event, event.ID)
			require.NotZero(t, participant.ID)
		}

	}
}
