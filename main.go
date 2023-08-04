package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"onursahin.dev/vuegolith/ui"

	"github.com/gorilla/mux"
)

type APIResponse struct {
	Ack  string      `json:"ack"`
	Data interface{} `json:"data"`
}

const VUEGOLITH_UPLOADS_DIR = "vuegolith-uploads"
const UPLOADS_ENDPOINT = "/uploads/"

func getHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir, nil
}

func createVuegolithUploadsDir() (string, error) {
	homeDir, err := getHomeDir()
	if err != nil {
		return "", err
	}

	uploadsDir := filepath.Join(homeDir, VUEGOLITH_UPLOADS_DIR)

	// Überprüfe, ob der Ordner bereits existiert
	_, err = os.Stat(uploadsDir)
	if os.IsNotExist(err) {
		// Ordner existiert nicht, also erstellen
		err = os.Mkdir(uploadsDir, 0755)
		if err != nil {
			return "", err
		}
	}

	return uploadsDir, nil
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

func handleLogGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Determine Path
	path, err := getHomeDir()
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Failed to determine home dir")
		return
	}

	file, err := os.Open(path + "/log.txt")
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Failed to read log")
		return
	}
	defer file.Close()

	// Read file
	c, err := io.ReadAll(file)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Failed to read log content")
		return
	}

	respondJSON(w, http.StatusOK, string(c[:]))
}

func handleLogPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Read body
	c, err := io.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Failed to read response body")
		return
	}
	defer r.Body.Close()

	// Determine Path
	path, err := getHomeDir()
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Failed to determine home dir")
		return
	}

	// Create or append to the log file
	file, err := os.OpenFile(path+"/log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Failed to create log file")
		return
	}
	defer file.Close()

	// Validate content
	content := string(c[:])
	if strings.TrimSpace(content) == "" {
		respondJSON(w, http.StatusInternalServerError, "Cannot save empty log")
		return
	}

	// Write the log entry to the file
	_, err = file.WriteString(string(c[:]) + "\n")
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Failed to write to log file")
		return
	}

	respondJSON(w, http.StatusOK, nil)
}

func handleListUploads(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Determine Path
	path, err := getHomeDir()
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Failed to determine home dir")
		return
	}

	uploadsDir := filepath.Join(path, VUEGOLITH_UPLOADS_DIR)

	// List all files in the uploads directory
	files, err := os.ReadDir(uploadsDir)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Failed to list uploads")
		return
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	respondJSON(w, http.StatusOK, fileNames)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Determine Path
	path, err := getHomeDir()
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Failed to determine home dir")
		return
	}

	err = r.ParseMultipartForm(10 << 30) // 1GB maximum file size
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Failed to parse form")
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, "Failed to get file from form")
		return
	}
	defer file.Close()

	// Save the uploaded file to the current directory
	f, err := os.OpenFile(filepath.Join(path+"/"+VUEGOLITH_UPLOADS_DIR, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Failed to save file")
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Failed to write file content")
		return
	}

	respondJSON(w, http.StatusOK, nil)
}

func main() {

	_, err := createVuegolithUploadsDir()
	if err != nil {
		panic(err)
	}

	assets, _ := ui.Assets()

	// Use the file system to serve static files
	fs := http.FileServer(http.FS(assets))
	http.Handle("/", http.StripPrefix("/", fs))

	// Create a new Router from Gorilla Mux
	router := mux.NewRouter()

	// Define API endpoints with Gorilla Mux
	router.HandleFunc("/api/log", handleLogGet).Methods("GET")
	router.HandleFunc("/api/log", handleLogPost).Methods("POST")
	router.HandleFunc("/api/uploads", handleUpload).Methods("POST")
	router.HandleFunc("/api/uploads", handleListUploads).Methods("GET")

	// Determine home dir
	homeDir, err := getHomeDir()
	if err != nil {
		panic(err)
	}
	fmt.Println(homeDir)

	// // Serve files from the home directory under the "/storage/" endpoint
	fs = http.FileServer(http.Dir(homeDir + "/" + VUEGOLITH_UPLOADS_DIR))
	http.Handle(UPLOADS_ENDPOINT, http.StripPrefix(UPLOADS_ENDPOINT, fs))

	corsHandler := corsMiddleware(router)
	http.Handle("/api/", corsHandler)

	// Pfade zu deinen selbstsignierten Zertifikatsdateien
	// port := "8484"

	//err = http.ListenAndServe(":"+port, nil)
	certFile := "server.crt"
	keyFile := "server.key"
	err = http.ListenAndServeTLS(":443", certFile, keyFile, nil)
	if err != nil {
		panic(err)
	}

}