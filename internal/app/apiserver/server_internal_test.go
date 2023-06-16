package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bruhlord-s/virttable-api/internal/app/model"
	"github.com/bruhlord-s/virttable-api/internal/app/store/teststore"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestServer_AuthenticateUser(t *testing.T) {
	store := teststore.New()
	u := model.TestUser(t)
	store.User().Create(u)

	jwtKey := "secret_key"

	token, _ := u.CreateJWT([]byte(jwtKey))
	fakeuser := u
	fakeuser.ID = uuid.New()
	fakeToken, _ := fakeuser.CreateJWT([]byte(jwtKey))
	
	testCases := []struct {
		name 	 	 string
		authHeader 	 string
		exceptedCode int
	} {
		{
			"authenticated",
			fmt.Sprintf("Bearer %s", token),
			http.StatusOK,
		},
		{
			"invalid jwt",
			"Bearer itsnotjwt",
			http.StatusBadRequest,
		},
		{
			"not authenticated",
			fmt.Sprintf("Bearer %s", fakeToken),
			http.StatusUnauthorized,
		},
		{
			"no auth header provided",
			"",
			http.StatusUnauthorized,
		},
		{
			"invalid auth header",
			"Bearer",
			http.StatusBadRequest,
		},
	}

	s := newServer(store, jwtKey)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("Authorization", tc.authHeader)
			s.authenticateUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, tc.exceptedCode, rec.Code)
		})
	}
}

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New(), "secret_key")

	testCases := []struct {
		name 		 string
		payload 	 interface{}
		exceptedCode int
	}{
		{
			"valid",
			map[string]string {
				"email": "test@test.com",
				"username": "test",
				"password": "password",
			},
			http.StatusCreated,
		},
		{
			"invalid payload",
			"some invalid payload",
			http.StatusBadRequest,
		},
		{
			"invalid body",
			map[string]string {
				"email": "wrong",
			},
			http.StatusUnprocessableEntity,
 		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.exceptedCode, rec.Code)
		})
	}
}

func TestServer_HandleSessionCreate(t *testing.T) {
	u := model.TestUser(t)
	store := teststore.New()
	store.User().Create(u)

	s := newServer(store, "secret_key")

	testCases := []struct {
		name 		 string
		payload 	 interface{}
		exceptedCode int
	} {
		{
			"valid",
			map[string]string {
				"email": u.Email,
				"password": u.Password,
			},
			http.StatusOK,
		},
		{
			"invalid payload",
			"some invalid payload",
			http.StatusBadRequest,
		},
		{
			"invalid email",
			map[string]string {
				"email": "wrong",
				"password": u.Password,
			},
			http.StatusUnauthorized,
 		},
		 {
			"invalid password",
			map[string]string {
				"email": u.Email,
				"password": "invalid",
			},
			http.StatusUnauthorized,
 		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.exceptedCode, rec.Code)
		})
	}
}