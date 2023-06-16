package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bruhlord-s/virttable-api/internal/app/model"
	"github.com/bruhlord-s/virttable-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

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