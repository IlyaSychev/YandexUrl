package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestFirstEndPoint(t *testing.T) {
	type want struct {
		code   int
		method string
		body   string
		header map[string][]string
	}
	tests := []struct {
		name    string
		request string
		reqBody string
		want    want
	}{
		{name: "1st_test", request: "/", reqBody: "https://practicum.yandex.ru/", want: want{code: http.StatusCreated, method: "POST", body: "http://localhost:8080/b", header: map[string][]string{"Content-Type": {"text/plain"}}}},
		{name: "2st_test", request: "/", reqBody: "https://practicum.yandex.ru/AS", want: want{code: http.StatusCreated, method: "POST", body: "http://localhost:8080/c", header: map[string][]string{"Content-Type": {"text/plain"}}}},
		{name: "3st_test", request: "/b", reqBody: "", want: want{code: 307, method: "GET", body: "", header: map[string][]string{"Location": {"https://practicum.yandex.ru/"}}}},
		{name: "4st_test", request: "/c", reqBody: "", want: want{code: 307, method: "GET", body: "", header: map[string][]string{"Location": {"https://practicum.yandex.ru/AS"}}}},
		{name: "5st_test", request: "/d", reqBody: "", want: want{code: 404, method: "GET", body: "", header: map[string][]string{"Location": {""}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.want.method == http.MethodPost {
				request := httptest.NewRequest(http.MethodPost, tt.request, strings.NewReader(tt.reqBody))
				w := httptest.NewRecorder()
				FirstEndPoint(w, request)

				res := w.Result()
				defer res.Body.Close()

				assert.Equal(t, tt.want.code, res.StatusCode)
				assert.Equal(t, tt.want.header["Content-Type"][0], res.Header.Get("Content-Type"))

				body, err := io.ReadAll(res.Body)
				assert.NoError(t, err)

				assert.Equal(t, tt.want.body, string(body))

			} else {
				r := mux.NewRouter()
				r.HandleFunc("/{id}", SecondEndPoint)
				request := httptest.NewRequest(http.MethodGet, tt.request, nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, request)
				res := w.Result()
				assert.Equal(t, tt.want.header["Location"][0], res.Header.Get("Location"))
				defer res.Body.Close()
				assert.Equal(t, tt.want.code, res.StatusCode)

			}

		})
	}
}
