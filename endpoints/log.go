package endpoints

import (
	"io"
	"net/http"
	"os"
	"strings"

	"onursahin.dev/vuegolith/api"
	"onursahin.dev/vuegolith/utils"
)

func HandleLogGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		api.RespondJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Determine Path
	path, err := utils.GetHomeDir()
	if err != nil {
		api.RespondJSON(w, http.StatusInternalServerError, "Failed to determine home dir")
		return
	}

	file, err := os.Open(path + "/log.txt")
	if err != nil {
		api.RespondJSON(w, http.StatusInternalServerError, "Failed to read log")
		return
	}
	defer file.Close()

	// Read file
	c, err := io.ReadAll(file)
	if err != nil {
		api.RespondJSON(w, http.StatusInternalServerError, "Failed to read log content")
		return
	}

	api.RespondJSON(w, http.StatusOK, string(c[:]))
}

func HandleLogPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		api.RespondJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Read body
	c, err := io.ReadAll(r.Body)
	if err != nil {
		api.RespondJSON(w, http.StatusInternalServerError, "Failed to read response body")
		return
	}
	defer r.Body.Close()

	// Determine Path
	path, err := utils.GetHomeDir()
	if err != nil {
		api.RespondJSON(w, http.StatusInternalServerError, "Failed to determine home dir")
		return
	}

	// Create or append to the log file
	file, err := os.OpenFile(path+"/log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		api.RespondJSON(w, http.StatusInternalServerError, "Failed to create log file")
		return
	}
	defer file.Close()

	// Validate content
	content := string(c[:])
	if strings.TrimSpace(content) == "" {
		api.RespondJSON(w, http.StatusInternalServerError, "Cannot save empty log")
		return
	}

	// Write the log entry to the file
	_, err = file.WriteString(string(c[:]) + "\n")
	if err != nil {
		api.RespondJSON(w, http.StatusInternalServerError, "Failed to write to log file")
		return
	}

	api.RespondJSON(w, http.StatusOK, nil)
}
