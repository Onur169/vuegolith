package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"onursahin.dev/vuegolith/ui"

	"github.com/gorilla/mux"
)

type APIResponse struct {
	Ack  string      `json:"ack"`
	Data interface{} `json:"data"`
}

func main() {
	fmt.Println("Hallo Welt!")
	assets, _ := ui.Assets()

	// Use the file system to serve static files
	fs := http.FileServer(http.FS(assets))
	http.Handle("/", http.StripPrefix("/", fs))

	// Create a new Router from Gorilla Mux
	router := mux.NewRouter()

	// Define API endpoints with Gorilla Mux
	router.HandleFunc("/api/log", handleLog).Methods("POST")
	router.HandleFunc("/api/upload", handleUpload).Methods("POST")

	// http.Handle("/api/", router)
	corsHandler := corsMiddleware(router)
	http.Handle("/api/", corsHandler)

	port := ":8080"
	fmt.Println("Server läuft auf http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
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

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	response := APIResponse{
		Ack:  "success",
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}

func handleLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read body
	c, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Create or append to the log file
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		http.Error(w, "Failed to create log file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fmt.Println(string(c))

	// Write the log entry to the file
	_, err = file.WriteString(string(c[:]) + "\n")
	if err != nil {
		http.Error(w, "Failed to write to log file", http.StatusInternalServerError)
		return
	}

	respondJSON(w, http.StatusOK, nil)
}


func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 30) // 1GB maximum file size
	if err != nil {
		fmt.Print(err.Error())
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the uploaded file to the current directory
	f, err := os.OpenFile(filepath.Join(".", handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Failed to write file content", http.StatusInternalServerError)
		return
	}

	respondJSON(w, http.StatusOK, nil)
}
