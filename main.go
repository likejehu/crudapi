package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/likejehu/crudapi/models"

	"github.com/gorilla/mux"
)

func readAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get(read) called"))
	fmt.Fprint(w, "\n Hello all")
	fmt.Fprint(w, "\n ", models.Books)
}

func read(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get(read) called"))
	fmt.Fprint(w, "\n Hello single")

}

func create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("post(create) called"))
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

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
	r.HandleFunc("/library", readAll).Methods(http.MethodGet)
	r.HandleFunc("/library", read).Methods(http.MethodGet)
	r.HandleFunc("/library", create).Methods(http.MethodPost)
	r.HandleFunc("/library", update).Methods(http.MethodPut)
	r.HandleFunc("/library", delete).Methods(http.MethodDelete)
	r.HandleFunc("/library", notFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
