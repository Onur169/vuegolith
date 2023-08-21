package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func FileServerWithCors(dir http.Dir) http.Handler {
	fs := http.FileServer(dir)
	return CorsMiddleware(fs)
}

func RouterWithCors(m *mux.Router) http.Handler {
	return CorsMiddleware(m)
}
