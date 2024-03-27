package db

import (
	"context"
	"testing"

	"github.com/TTKirito/backend-go/utils"
	"github.com/stretchr/testify/require"
)

func createRandomLocation(t *testing.T) Location {
	event := createRandomEvent(t)

	arg := CreateLocationParams{
		Lat:    utils.RandomLatLong().Lat,
		Long:   utils.RandomLatLong().Long,
		Street: utils.RandomLatLong().Street,
		Event:  event.ID,
	}

	location, err := testQueries.CreateLocation(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, location.Lat, arg.Lat)
	require.Equal(t, location.Long, arg.Long)
	require.Equal(t, location.Street, arg.Street)

	require.NotZero(t, location.ID)
	require.NotZero(t, location.CreatedAt)

	return location
}

func TestCreateLocation(t *testing.T) {
	createRandomLocation(t)
}

func TestGetLocation(t *testing.T) {
	location := createRandomLocation(t)
	location1, err := testQueries.GetLocation(context.Background(), location.Event)
	require.NoError(t, err)
	require.Equal(t, location.Lat, location1.Lat)
	require.Equal(t, location.Long, location1.Long)
	require.Equal(t, location.Street, location1.Street)
	require.Equal(t, location.ID, location1.ID)

	require.NotZero(t, location.ID)
	require.NotZero(t, location.CreatedAt)
}

func TestUpdateLocation(t *testing.T) {
	location := createRandomLocation(t)
	arg := UpdateLocationParams{
		Lat:    utils.RandomLatLong().Lat,
		Long:   utils.RandomLatLong().Long,
		Street: utils.RandomLatLong().Street,
		ID:     location.ID,
	}
	location1, err := testQueries.UpdateLocation(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Lat, location1.Lat)
	require.Equal(t, arg.Long, location1.Long)
	require.Equal(t, arg.Street, location1.Street)
	require.Equal(t, location.ID, location1.ID)

	require.NotZero(t, location.ID)
	require.NotZero(t, location.CreatedAt)
}
