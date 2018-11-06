package handler

import (
  "log"
  "net/http"
  "encoding/json"

  "github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"vblog/pkg/entity"
  "vblog/pkg/auth"
)

func login(service auth.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    errorMessage := "Login error"
    var u *entity.User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
    auth, err := service.Login(u)
    if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

    w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(auth); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
  })
}

func logout(service auth.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    errorMessage := "Logout error"
    var auth *entity.Auth
		err := json.NewDecoder(r.Body).Decode(&auth)
    if err != nil {
      log.Println("decode error")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		err = service.Logout(auth.ID)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
      log.Println("logout error")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
  })
}

//-----------------------------------------------------------------------------
func MakeAuthHandler(r *mux.Router, n negroni.Negroni, service auth.Repository) {
  r.Handle("/v1/auth/login", n.With(
		negroni.Wrap(login(service)),
	)).Methods("POST", "OPTIONS").Name("login")

	r.Handle("/v1/auth/logout", n.With(
		negroni.Wrap(logout(service)),
	)).Methods("POST", "OPTIONS").Name("logout")
}
