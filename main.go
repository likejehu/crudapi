package main

import (
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "welcome to library"}`))
}

func main() {
	s := &server{}
	http.Handle("/library", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
