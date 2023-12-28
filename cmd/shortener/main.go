package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/IlyacSychev/YandexUrl/internal/app"
	"github.com/gorilla/mux"
)

var (
	urlMap     = make(map[string]string)
	urlMapLock sync.RWMutex
)

func FirstEndPoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)

	shortURL := app.Short(string(body))

	urlMapLock.Lock()
	urlMap[shortURL] = string(body)
	urlMapLock.Unlock()

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "http://localhost:8080/"+shortURL)
}

func SecondEndPoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["id"]
	urlMapLock.RLock()
	originalURL, ok := urlMap[name]
	urlMapLock.RUnlock()
	if ok {
		http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", FirstEndPoint)
	mux.HandleFunc("/{id}", SecondEndPoint)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
