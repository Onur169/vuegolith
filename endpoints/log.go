package endpoints

import (
	"io"
	"net/http"
	"os"
	"strings"

	"onursahin.dev/vuegolith/api/response"
	"onursahin.dev/vuegolith/utils"
)

func HandleLogGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.RespondJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	path, err := utils.GetHomeDir()
	if err != nil {
		response.RespondJSON(w, http.StatusInternalServerError, "Failed to determine home dir")
		return
	}

	file, err := os.Open(path + "/log.txt")
	if err != nil {
		response.RespondJSON(w, http.StatusInternalServerError, "Failed to read log")
		return
	}
	defer file.Close()

	// Read file
	c, err := io.ReadAll(file)
	if err != nil {
		response.RespondJSON(w, http.StatusInternalServerError, "Failed to read log content")
		return
	}

	response.RespondJSON(w, http.StatusOK, string(c[:]))
}

func HandleLogPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.RespondJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	c, err := io.ReadAll(r.Body)
	if err != nil {
		response.RespondJSON(w, http.StatusInternalServerError, "Failed to read response body")
		return
	}
	defer r.Body.Close()

	path, err := utils.GetHomeDir()
	if err != nil {
		response.RespondJSON(w, http.StatusInternalServerError, "Failed to determine home dir")
		return
	}

	content := string(c[:])
	if strings.TrimSpace(content) == "" {
		response.RespondJSON(w, http.StatusInternalServerError, "Cannot save empty log")
		return
	}

	file, err := os.OpenFile(path+"/log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		response.RespondJSON(w, http.StatusInternalServerError, "Failed to create log file")
		return
	}
	defer file.Close()

	_, err = file.WriteString(content + "\n")
	if err != nil {
		response.RespondJSON(w, http.StatusInternalServerError, "Failed to write to log file")
		return
	}

	response.RespondJSON(w, http.StatusOK, nil)
}
