package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bruhlord-s/virttable-api/internal/app/model"
	"github.com/bruhlord-s/virttable-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	ErrIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

type server struct {
	router 	 *mux.Router
	logger 	 *logrus.Logger
	store 	 store.Store
	jwtKey	 string
}

func newServer(store store.Store, jwtKey string) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store: store,
		jwtKey: jwtKey,
	}

	s.configureRouter()
	s.logger.Info("starting API server")

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email 	 string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("request to /users")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email: req.Email,
			Username: req.Username,
			Password: req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {
	type request struct {
		Email 	 string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("request to /sessions")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, ErrIncorrectEmailOrPassword)
			return
		}

		t, err := u.CreateJWT([]byte(s.jwtKey))
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, nil)
			return
		}
		response := map[string]string{
			"token": t,
		}
		s.respond(w, r, http.StatusOK, response)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
	s.logger.Error(err.Error())
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}