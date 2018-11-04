package main

import (
  "os"
  "time"
  "log"
  "net/http"
  "github.com/gorilla/context"
  "github.com/gorilla/mux"
)

const PORT = "9009"

func main()  {
  l := log.Println

  // start db connection session

  // route config
  r := mux.NewRouter()

  // handle and start sersion
  http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

  logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + PORT,
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}


  // log error if raised

  l("API server is running")
}
