package api

import "net/http"

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins for demonstration purposes.
		// In a production environment, you might want to restrict this to specific origins.
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Set allowed headers and methods.
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			// Handle preflight requests by responding with 200 OK.
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler.
		next.ServeHTTP(w, r)
	})
}

func FileServerWithCors(dir http.Dir) http.Handler {
	fs := http.FileServer(dir)
	return CorsMiddleware(fs)
}