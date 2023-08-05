package endpoints

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"onursahin.dev/vuegolith/api"
	"onursahin.dev/vuegolith/utils"
)

type UploadsList struct {
	Name string `json:"name"`
}

func HandleListUploads(w http.ResponseWriter, r *http.Request) {
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

	uploadsDir := filepath.Join(path, utils.GetUploadsDirName())

	// List all files in the uploads directory
	files, err := os.ReadDir(uploadsDir)
	if err != nil {
		api.RespondJSON(w, http.StatusInternalServerError, "Failed to list uploads")
		return
	}

	var fileNames []UploadsList
	for _, file := range files {
		fileNames = append(fileNames, UploadsList{Name: file.Name()})
	}

	api.RespondJSON(w, http.StatusOK, fileNames)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		api.RespondJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Determine Path
	path, err := utils.GetHomeDir()
	if err != nil {
		api.RespondJSON(w, http.StatusInternalServerError, "Failed to determine home dir")
		return
	}

	err = r.ParseMultipartForm(10 << 30) // 1GB maximum file size
	if err != nil {
		api.RespondJSON(w, http.StatusInternalServerError, "Failed to parse form")
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		api.RespondJSON(w, http.StatusBadRequest, "Failed to get file from form")
		return
	}
	defer file.Close()

	// Save the uploaded file to the current directory
	f, err := os.OpenFile(filepath.Join(path+"/"+utils.GetUploadsDirName(), handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		api.RespondJSON(w, http.StatusInternalServerError, "Failed to save file")
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		api.RespondJSON(w, http.StatusInternalServerError, "Failed to write file content")
		return
	}

	api.RespondJSON(w, http.StatusOK, nil)
}
