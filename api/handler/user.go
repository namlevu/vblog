package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"vblog/pkg/entity"
	"vblog/pkg/user"
)

func userIndex(service user.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("user Index called")
		errorMessage := "Error reading bookmarks"
		var data []*entity.User
		var err error
		query := r.URL.Query().Get("query")
		switch {
		case query == "":
			log.Println("user Index query empty")
			data, err = service.SelectAll()
		default:
			log.Println("user Index query exist")
			data, err = service.Search(query)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func userInsert(service user.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding bookmark"
		var u *entity.User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		u.ID, err = service.Insert(u)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(u); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func MakeUserHandler(r *mux.Router, n negroni.Negroni, service user.Repository) {
	r.Handle("/v1/users", n.With(
		negroni.Wrap(userIndex(service)),
	)).Methods("GET", "OPTIONS").Name("userIndex")

	r.Handle("/v1/users", n.With(
		negroni.Wrap(userInsert(service)),
	)).Methods("POST", "OPTIONS").Name("userInsert")
}
