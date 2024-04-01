package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	mockdb "github.com/TTKirito/backend-go/db/mock"
	db "github.com/TTKirito/backend-go/db/sqlc"
	"github.com/TTKirito/backend-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := utils.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func TestCreateUser(t *testing.T) {
	user, password := randomUser(t)

	ctrl := gomock.NewController(t)
	store := mockdb.NewMockStore(ctrl)

	body := gin.H{
		"username":  user.Username,
		"full_name": user.FullName,
		"password":  password,
		"email":     user.Email,
	}

	arg := db.CreateUserParams{
		Username:       user.Username,
		FullName:       user.FullName,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
	}

	store.EXPECT().
		CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
		Times(1).
		Return(user, nil)

	server := newTestServer(t, store)
	recorder := httptest.NewRecorder()

	data, err := json.Marshal(body)
	require.NoError(t, err)

	url := "/users"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	require.NoError(t, err)

	server.route.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchUser(t, recorder.Body, user)
}

func TestLoginUser(t *testing.T) {
	user, password := randomUser(t)

	testCase := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username": user.Username,
				"password": password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUser(gomock.Any(), user.Username).Times(1).Return(user, nil)

				store.EXPECT().
					CreateSession(gomock.Any(), gomock.Any()).
					Times(1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

			},
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			store := mockdb.NewMockStore(ctrl)

			tc.buildStubs(store)
			server := newTestServer(t, store)

			recorder := httptest.NewRecorder()

			url := "/users/login"
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.route.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func randomUser(t *testing.T) (user db.User, password string) {
	password = utils.RandomString(6)

	hashedPassword, err := utils.HashedPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username:          utils.RandomOwner(),
		FullName:          utils.RandomOwner(),
		HashedPassword:    hashedPassword,
		Email:             utils.RandomEmail(),
		CreatedAt:         utils.RandomTime(),
		PasswordChangedAt: utils.RandomTime(),
	}
	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)

	require.NoError(t, err)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.FullName, gotUser.FullName)
	require.Equal(t, user.Email, gotUser.Email)
	require.Empty(t, gotUser.HashedPassword)
}
