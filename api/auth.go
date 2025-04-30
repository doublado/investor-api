package api

import (
	"net/http"
	"os"
)

const headerName = "X-API-SECRET"

func RequireAuth(handler http.HandlerFunc) http.HandlerFunc {
	expected := os.Getenv("API_SECRET")

	return func(w http.ResponseWriter, r *http.Request) {
		provided := r.Header.Get(headerName)
		if provided != expected {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		handler(w, r)
	}
}
