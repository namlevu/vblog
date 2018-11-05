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

//-----------------------------------------------------------------------------
// handler functions
func postIndex(service post.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println("Post Index called")
		errorMessage := "Error reading posts"
		var data []*entity.Post
		var err error
		query := r.URL.Query().Get("query")
		switch {
		case query == "":
			log.Println("Post Index query empty")
			data, err = service.SelectAll()
		default:
			log.Println("Post Index query exist")
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
func postInsert(service post.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    errorMessage := "Error adding post"
		var p *entity.Post
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		p.ID, err = service.Insert(p)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(p); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
  })
}
//-----------------------------------------------------------------------------
func MakePostHandler(r *mux.Router, n negroni.Negroni, service post.Repository) {
  r.Handle("/v1/posts", n.With(
		negroni.Wrap(postIndex(service)),
	)).Methods("GET", "OPTIONS").Name("postIndex")

	r.Handle("/v1/posts", n.With(
		negroni.Wrap(postInsert(service)),
	)).Methods("POST", "OPTIONS").Name("postInsert")
}
