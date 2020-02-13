package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func read(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get(read) called"))
}

func create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("post(create) called"))
}

func update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("put(update) called"))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("delete called"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/library", read).Methods(http.MethodGet)
	r.HandleFunc("/library", create).Methods(http.MethodPost)
	r.HandleFunc("/library", update).Methods(http.MethodPut)
	r.HandleFunc("/library", delete).Methods(http.MethodDelete)
	r.HandleFunc("/library", notFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
