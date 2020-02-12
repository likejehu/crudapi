package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/library", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Игорь, добро пожаловать в библиотеку")
	})
	http.ListenAndServe(":8080", nil)
}
