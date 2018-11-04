package main

import (
  "os"
  "time"
  "log"
  "net/http"
  "github.com/gorilla/context"
  "github.com/gorilla/mux"
)

const PORT = ":9009"
const TIMEOUT_READ = 5 * time.Second
const TIMEOUT_WRITE = 10 * time.Second

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
		ReadTimeout: TIMEOUT_READ,
		WriteTimeout: TIMEOUT_WRITE,
		Addr: PORT,
		Handler: context.ClearHandler(http.DefaultServeMux),
		ErrorLog: logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}


  // log error if raised

  l("API server is running")
}
