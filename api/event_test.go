package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/TTKirito/backend-go/db/mock"
	db "github.com/TTKirito/backend-go/db/sqlc"
	"github.com/TTKirito/backend-go/token"
	"github.com/TTKirito/backend-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateEventAPI(t *testing.T) {
	user, _ := randomUser(t)
	owner := randomAccount(user.Username)
	account := randomAccount(user.Username)
	event := randomEvent(owner)
	participants := randomParticipant(account, event)
	location := randomLocation(event)
	eventTx := randomEventTx(event, participants, location)

	// body := gin.H{
	// 	"title":        eventTx.Event.Title.String,
	// 	"start_time":   eventTx.Event.StartTime,
	// 	"end_time":     eventTx.Event.EndTime,
	// 	"is_emegency":  eventTx.Event.IsEmegency,
	// 	"owner":        eventTx.Event.Owner,
	// 	"note":         eventTx.Event.Note.String,
	// 	"type":         eventTx.Event.Type,
	// 	"visit_type":   eventTx.Event.VisitType,
	// 	"meeting":      eventTx.Event.Meeting.String,
	// 	"participants": eventTx.Participants,
	// 	"location":     eventTx.Location,
	// }

	// // mock store
	// ctrl := gomock.NewController(t)
	// store := mockdb.NewMockStore(ctrl)

	// // build stubs

	// arg := db.CreateEventTxParams{
	// 	Title:        eventTx.Event.Title.String,
	// 	StartTime:    eventTx.Event.StartTime,
	// 	EndTime:      eventTx.Event.EndTime,
	// 	IsEmegency:   eventTx.Event.IsEmegency,
	// 	Owner:        eventTx.Event.Owner,
	// 	Note:         eventTx.Event.Note.String,
	// 	Type:         eventTx.Event.Type,
	// 	VisitType:    eventTx.Event.VisitType,
	// 	Meeting:      eventTx.Event.Meeting.String,
	// 	Participants: eventTx.Participants,
	// 	Location:     eventTx.Location,
	// }

	// store.EXPECT().CreateEventTx(gomock.Any(), gomock.Eq(arg)).Times(1).Return(eventTx, nil)

	// // start new server and test request

	// // require.Equal(t, http.StatusOK, recorder.Code)
	// // requireBodyMatchAccount(t, recorder.Body, account)

	// server := newTestServer(t, store)
	// recorder := httptest.NewRecorder()
	// url := "/events"

	// data, err := json.Marshal(body)
	// require.NoError(t, err)

	// request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	// require.NoError(t, err)

	// server.route.ServeHTTP(recorder, request)
	// // check response

	// require.Equal(t, http.StatusOK, recorder.Code)
	// requireBodyMatchEventTx(t, recorder.Body, eventTx)

	testCase := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"title":        eventTx.Event.Title.String,
				"start_time":   eventTx.Event.StartTime,
				"end_time":     eventTx.Event.EndTime,
				"is_emegency":  eventTx.Event.IsEmegency,
				"owner":        eventTx.Event.Owner,
				"note":         eventTx.Event.Note.String,
				"type":         eventTx.Event.Type,
				"visit_type":   eventTx.Event.VisitType,
				"meeting":      eventTx.Event.Meeting.String,
				"participants": eventTx.Participants,
				"location":     eventTx.Location,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)

			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateEventTxParams{
					Title:        eventTx.Event.Title.String,
					StartTime:    eventTx.Event.StartTime,
					EndTime:      eventTx.Event.EndTime,
					IsEmegency:   eventTx.Event.IsEmegency,
					Owner:        eventTx.Event.Owner,
					Note:         eventTx.Event.Note.String,
					Type:         eventTx.Event.Type,
					VisitType:    eventTx.Event.VisitType,
					Meeting:      eventTx.Event.Meeting.String,
					Participants: eventTx.Participants,
					Location:     eventTx.Location,
				}
				store.EXPECT().CreateEventTx(gomock.Any(), arg).Times(1).Return(eventTx, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchEventTx(t, recorder.Body, eventTx)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"title":        eventTx.Event.Title.String,
				"start_time":   eventTx.Event.StartTime,
				"end_time":     eventTx.Event.EndTime,
				"is_emegency":  eventTx.Event.IsEmegency,
				"owner":        eventTx.Event.Owner,
				"note":         eventTx.Event.Note.String,
				"type":         eventTx.Event.Type,
				"visit_type":   eventTx.Event.VisitType,
				"meeting":      eventTx.Event.Meeting.String,
				"participants": eventTx.Participants,
				"location":     eventTx.Location,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)

			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateEventTxParams{
					Title:        eventTx.Event.Title.String,
					StartTime:    eventTx.Event.StartTime,
					EndTime:      eventTx.Event.EndTime,
					IsEmegency:   eventTx.Event.IsEmegency,
					Owner:        eventTx.Event.Owner,
					Note:         eventTx.Event.Note.String,
					Type:         eventTx.Event.Type,
					VisitType:    eventTx.Event.VisitType,
					Meeting:      eventTx.Event.Meeting.String,
					Participants: eventTx.Participants,
					Location:     eventTx.Location,
				}
				store.EXPECT().CreateEventTx(gomock.Any(), arg).Times(1).Return(db.CreateEventTxResult{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidType",
			body: gin.H{
				"title":        eventTx.Event.Title.String,
				"start_time":   eventTx.Event.StartTime,
				"end_time":     eventTx.Event.EndTime,
				"is_emegency":  eventTx.Event.IsEmegency,
				"owner":        eventTx.Event.Owner,
				"note":         eventTx.Event.Note.String,
				"type":         "account",
				"visit_type":   eventTx.Event.VisitType,
				"meeting":      eventTx.Event.Meeting.String,
				"participants": eventTx.Participants,
				"location":     eventTx.Location,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)

			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateEventTxParams{
					Title:        eventTx.Event.Title.String,
					StartTime:    eventTx.Event.StartTime,
					EndTime:      eventTx.Event.EndTime,
					IsEmegency:   eventTx.Event.IsEmegency,
					Owner:        eventTx.Event.Owner,
					Note:         eventTx.Event.Note.String,
					Type:         eventTx.Event.Type,
					VisitType:    eventTx.Event.VisitType,
					Meeting:      eventTx.Event.Meeting.String,
					Participants: eventTx.Participants,
					Location:     eventTx.Location,
				}
				store.EXPECT().CreateEventTx(gomock.Any(), arg).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]
		ctrl := gomock.NewController(t)
		store := mockdb.NewMockStore(ctrl)
		tc.buildStubs(store)

		server := newTestServer(t, store)
		recorder := httptest.NewRecorder()
		url := "/events"

		data, err := json.Marshal(tc.body)
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		require.NoError(t, err)
		tc.setupAuth(t, request, server.tokenMaker)

		server.route.ServeHTTP(recorder, request)
		tc.checkResponse(t, recorder)
	}
}

func TestGetEventAPI(t *testing.T) {
	user, _ := randomUser(t)

	owner := randomAccount(user.Username)
	event := randomEvent(owner)

	testCase := []struct {
		name          string
		eventID       int64
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			eventID: event.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)

			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetEvent(gomock.Any(), gomock.Eq(event.ID)).Times(1).Return(event, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchEvent(t, recorder.Body, event)
			},
		},
		{
			name:    "NotFound",
			eventID: event.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)

			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetEvent(gomock.Any(), gomock.Any()).Times(1).Return(db.Event{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:    "InternalError",
			eventID: event.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)

			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetEvent(gomock.Any(), gomock.Any()).Times(1).Return(db.Event{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:    "InvalidID",
			eventID: 0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)

			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetEvent(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]
		ctrl := gomock.NewController(t)
		store := mockdb.NewMockStore(ctrl)
		tc.buildStubs(store)

		server := newTestServer(t, store)
		recorder := httptest.NewRecorder()

		url := fmt.Sprintf("/events/%d", tc.eventID)

		request, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)
		tc.setupAuth(t, request, server.tokenMaker)
		server.route.ServeHTTP(recorder, request)
		tc.checkResponse(t, recorder)
	}

}

func TestListEventAPI(t *testing.T) {
	user, _ := randomUser(t)

	n := 5
	events := make([]db.Event, n)
	for i := 0; i < n; i++ {
		events[i] = randomEvent(randomAccount(user.Username))
	}

	type Query struct {
		PageID    int
		PageSize  int
		StartTime int64
		EndTime   int64
	}

	testCase := []struct {
		name          string
		Query         Query
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			Query: Query{
				PageID:    1,
				PageSize:  n,
				StartTime: time.Now().UTC().AddDate(0, 0, -1).Unix(),
				EndTime:   time.Now().UTC().Add(30 * time.Minute).Unix(),
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)

			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListEventParams{
					Limit:     int32(n),
					Offset:    0,
					StartTime: time.Now().UTC().AddDate(0, 0, -1).Unix(),
					EndTime:   time.Now().UTC().Add(30 * time.Minute).Unix(),
				}
				store.EXPECT().ListEvent(gomock.Any(), gomock.Eq(arg)).Times(1).Return(events, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchEvents(t, recorder.Body, events)
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := mockdb.NewMockStore(ctrl)
		tc.buildStubs(store)

		server := newTestServer(t, store)
		recorder := httptest.NewRecorder()

		url := "/events"

		request, err := http.NewRequest(http.MethodGet, url, nil)

		require.NoError(t, err)
		q := request.URL.Query()
		q.Add("page_id", fmt.Sprintf("%d", tc.Query.PageID))
		q.Add("page_size", fmt.Sprintf("%d", tc.Query.PageSize))
		q.Add("start_time", fmt.Sprintf("%d", tc.Query.StartTime))
		q.Add("end_time", fmt.Sprintf("%d", tc.Query.EndTime))
		request.URL.RawQuery = q.Encode()
		tc.setupAuth(t, request, server.tokenMaker)

		server.route.ServeHTTP(recorder, request)
		tc.checkResponse(t, recorder)

	}
}

func randomEvent(account db.Account) db.Event {
	return db.Event{
		ID:         utils.RandomInt(1, 10000),
		Title:      sql.NullString{String: utils.RandomString(12), Valid: true},
		StartTime:  utils.RandomTime().Unix(),
		EndTime:    utils.RandomTime().Add(30 * time.Minute).Unix(),
		IsEmegency: utils.RandomEmegency(),
		Owner:      account.ID,
		Note:       sql.NullString{String: utils.RandomString(12), Valid: true},
		Type:       db.EventTypes(utils.RandomEventType()),
		VisitType:  db.VisitTypes(utils.RandomVisitType()),
		Meeting:    sql.NullString{String: utils.RandomString(12), Valid: true},
	}
}

func randomParticipant(account db.Account, event db.Event) []db.Participant {
	return []db.Participant{
		{
			ID: utils.RandomInt(1, 10000),

			Event:   event.ID,
			Account: account.ID,
		},
	}
}

func randomLocation(event db.Event) db.Location {
	return db.Location{
		ID:     utils.RandomInt(1, 10000),
		Lat:    utils.RandomLatLong().Lat,
		Long:   utils.RandomLatLong().Long,
		Street: utils.RandomLatLong().Street,
		Event:  event.ID,
	}
}

func randomEventTx(event db.Event, participants []db.Participant, location db.Location) db.CreateEventTxResult {
	return db.CreateEventTxResult{
		Event:        event,
		Participants: participants,
		Location:     location,
	}
}

func requireBodyMatchEventTx(t *testing.T, body *bytes.Buffer, eventTx db.CreateEventTxResult) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)
	var gotEventTx db.CreateEventTxResult
	err = json.Unmarshal(data, &gotEventTx)
	require.NoError(t, err)
	require.Equal(t, eventTx, gotEventTx)
}

func requireBodyMatchEvent(t *testing.T, body *bytes.Buffer, event db.Event) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)
	var gotEvent db.Event
	err = json.Unmarshal(data, &gotEvent)
	require.NoError(t, err)
	require.Equal(t, gotEvent, event)
}

func requireBodyMatchEvents(t *testing.T, body *bytes.Buffer, event []db.Event) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)
	var gotEvents []db.Event
	err = json.Unmarshal(data, &gotEvents)
	require.NoError(t, err)
	require.Equal(t, gotEvents, event)

}
