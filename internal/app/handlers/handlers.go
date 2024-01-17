package handlers

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

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	if r.Method != http.MethodPost {
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	shortURL := app.Short(string(body))

	urlMapLock.Lock()
	urlMap[shortURL] = string(body)
	urlMapLock.Unlock()

	fmt.Fprint(w, "http://localhost:8080/"+shortURL)
}

func SecondEndPoint(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["id"]
	urlMapLock.RLock()
	originalURL, ok := urlMap[name]
	urlMapLock.RUnlock()
	if ok {
		w.Header().Set("Location", originalURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	http.NotFound(w, r)

}
