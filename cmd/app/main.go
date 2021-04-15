package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	var (
		addr = "0.0.0.0:8080"
		prefix = "/test/api"
	)

	router := mux.NewRouter()
	router.HandleFunc(prefix, indexHandle).Methods("GET")
	router.HandleFunc(prefix+"/ping", pingHandle).Methods("GET")

	server := http.Server{Addr: addr, Handler: router}
	log.Print(addr+prefix + " serving")

	panic(server.ListenAndServe())
}

func indexHandle(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(203)
	w.Write([]byte("index page"))
}

func pingHandle(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(200)
	w.Write([]byte("Pong!"))
}

