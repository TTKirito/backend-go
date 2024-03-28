package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/TTKirito/backend-go/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEvent(t *testing.T) Event {
	account := createRandomAccount(t)
	arg := CreateEventParams{
		Title:      sql.NullString{String: utils.RandomString(12), Valid: true},
		StartTime:  utils.RandomTime().Unix(),
		EndTime:    utils.RandomTime().Add(30 * time.Minute).Unix(),
		IsEmegency: utils.RandomEmegency(),
		Owner:      account.ID,
		Note:       sql.NullString{String: utils.RandomString(12), Valid: true},
		Type:       EventTypes(utils.RandomEventType()),
		VisitType:  VisitTypes(utils.RandomVisitType()),
		Meeting:    sql.NullString{String: "http://", Valid: true},
	}

	event, err := testQueries.CreateEvent(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Title, event.Title)
	require.Equal(t, arg.StartTime, event.StartTime)
	require.Equal(t, arg.EndTime, event.EndTime)
	require.Equal(t, arg.IsEmegency, event.IsEmegency)
	require.Equal(t, arg.Owner, event.Owner)
	require.Equal(t, arg.Note, event.Note)
	require.Equal(t, arg.Type, event.Type)
	require.Equal(t, arg.VisitType, event.VisitType)
	require.Equal(t, arg.Meeting, event.Meeting)

	require.NotZero(t, event.CreatedAt)
	require.NotZero(t, event.ID)

	return event
}

func TestCreateEvent(t *testing.T) {
	createRandomEvent(t)
}

func TestGetEvent(t *testing.T) {
	event := createRandomEvent(t)

	event1, err := testQueries.GetEvent(context.Background(), event.ID)

	require.NoError(t, err)
	require.Equal(t, event1.Title, event.Title)
	require.Equal(t, event1.StartTime, event.StartTime)
	require.Equal(t, event1.EndTime, event.EndTime)
	require.Equal(t, event1.IsEmegency, event.IsEmegency)
	require.Equal(t, event1.Owner, event.Owner)
	require.Equal(t, event1.Note, event.Note)
	require.Equal(t, event1.Type, event.Type)
	require.Equal(t, event1.VisitType, event.VisitType)
	require.Equal(t, event1.Meeting, event.Meeting)

	require.NotZero(t, event1.CreatedAt)
	require.NotZero(t, event1.ID)
}

func TestUpdateEvent(t *testing.T) {
	event := createRandomEvent(t)
	account := createRandomAccount(t)

	arg := UpdateEventParams{
		ID:         event.ID,
		Title:      sql.NullString{String: utils.RandomString(12), Valid: true},
		StartTime:  utils.RandomTime().Unix(),
		EndTime:    utils.RandomTime().Add(30 * time.Minute).Unix(),
		IsEmegency: utils.RandomEmegency(),
		Owner:      account.ID,
		Note:       sql.NullString{String: utils.RandomString(12), Valid: true},
		Type:       EventTypes(utils.RandomEventType()),
		VisitType:  VisitTypes(utils.RandomVisitType()),
		Meeting:    sql.NullString{String: "http://", Valid: true},
	}

	event1, err := testQueries.UpdateEvent(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, event1.Title, arg.Title)
	require.Equal(t, event1.StartTime, arg.StartTime)
	require.Equal(t, event1.EndTime, arg.EndTime)
	require.Equal(t, event1.IsEmegency, arg.IsEmegency)
	require.Equal(t, event1.Owner, arg.Owner)
	require.Equal(t, event1.Note, arg.Note)
	require.Equal(t, event1.Type, arg.Type)
	require.Equal(t, event1.VisitType, arg.VisitType)
	require.Equal(t, event1.Meeting, arg.Meeting)
	require.Equal(t, event1.ID, arg.ID)

	require.NotZero(t, event1.CreatedAt)
	require.NotZero(t, event1.ID)
}

func TestDeleteEvent(t *testing.T) {
	event := createRandomEvent(t)

	err := testQueries.DeleteEvent(context.Background(), event.ID)

	require.NoError(t, err)

	event1, err := testQueries.GetEvent(context.Background(), event.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, event1)
}

func TestListEvent(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEvent(t)
	}

	arg := ListEventParams{
		StartTime: time.Now().UTC().AddDate(0, 0, -1).Unix(),
		EndTime:   time.Now().UTC().Add(30 * time.Minute).Unix(),
		Limit:     5,
		Offset:    0,
	}

	events, err := testQueries.ListEvent(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, events, 5)

	for _, event := range events {
		require.NotEmpty(t, event)
	}

}
