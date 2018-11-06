package main

import (
	"log"
	"net/http"
	"os"
	"time"
  "strconv"

  "github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"

  "vblog/api/handler"
  "vblog/cnf"
	"vblog/pkg/auth"
  "vblog/pkg/user"
	"vblog/pkg/post"
  "vblog/pkg/middleware"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	l := log.Println
	w.WriteHeader(http.StatusOK)
	l("Client pinged!", r.Header.Get("ClientId"))
	return
}

func main() {
	l := log.Println
	l("API server is starting")
	var configuration cnf.Configuration
	err := cnf.LoadConfig(&configuration, cnf.DEV)
	if err != nil {
		log.Fatal(err.Error())
	}
	// start db connection session
	session, err := mgo.Dial(configuration.MongodbHost)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	mPool := mgosession.NewPool(nil, session, configuration.MongodbConnectionPool)
	defer mPool.Close()
	//
	userRepo := user.NewRepositoryMongo(configuration.MongodbDatabase, mPool)
	userService := user.NewService(userRepo)

	// route config
	r := mux.NewRouter()

	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)
	handler.MakeUserHandler(r, *n, userService)

	postRepo := post.NewRepositoryMongo(configuration.MongodbDatabase, mPool)
	postService := post.NewService(postRepo)
	handler.MakePostHandler(r, *n, postService)

	authRepo := auth.NewRepositoryMongo(configuration.MongodbDatabase, mPool)
	authService := auth.NewService(authRepo)
	handler.MakeAuthHandler(r, *n, authService)

	// handle and start sersion
	http.Handle("/", r)
	r.HandleFunc("/ping", pingHandler)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(configuration.ApiPort),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	// log error if raised
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
