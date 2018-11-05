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
  "vblog/cnf/development"
  "vblog/pkg/user"
  "vblog/pkg/middleware"
)

const PORT = ":9009"
const TIMEOUT_READ = 5 * time.Second
const TIMEOUT_WRITE = 10 * time.Second

func pingHandler(w http.ResponseWriter, r *http.Request) {
	l := log.Println
	w.WriteHeader(http.StatusOK)
	l("Client pinged!", r.Header.Get("ClientId"))
	return
}

func main() {
	l := log.Println
	// start db connection session
	session, err := mgo.Dial(development.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	mPool := mgosession.NewPool(nil, session, development.MONGODB_CONNECTION_POOL)
	defer mPool.Close()
	//
	userRepo := user.NewRepositoryMongo(development.MONGODB_DATABASE, mPool)
	userService := user.NewService(userRepo)

	// route config
	r := mux.NewRouter()

	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)
	handler.MakeUserHandler(r, *n, userService)
	// handle and start sersion
	http.Handle("/", r)
	r.HandleFunc("/ping", pingHandler)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  TIMEOUT_READ,
		WriteTimeout: TIMEOUT_WRITE,
		Addr:         ":" + strconv.Itoa(development.API_PORT),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	// log error if raised
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

	l("API server is running")
}
