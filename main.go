package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	port := "8080"

	r := chi.NewRouter()

	r.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Connected to server instance so networking is working")
	}))
	r.Handle("/helloWorld", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello world")
	}))

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
