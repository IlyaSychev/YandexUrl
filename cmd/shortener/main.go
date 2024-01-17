package main

import (
	"net/http"

	"github.com/IlyacSychev/YandexUrl/internal/app/handlers"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", handlers.FirstEndPoint)
	mux.HandleFunc("/{id}", handlers.SecondEndPoint)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
