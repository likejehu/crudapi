package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func road(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("get(read) called"))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("post(create) called"))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("put(update) called"))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("delete called"))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
	}

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/library", road)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
