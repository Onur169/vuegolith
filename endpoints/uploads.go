package endpoints

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"onursahin.dev/vuegolith/api/response"
	"onursahin.dev/vuegolith/utils"
)

type UploadsList struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

func HandleListUploads(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WithJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	path, err := utils.GetHomeDir()
	if err != nil {
		response.WithJSON(w, http.StatusInternalServerError, "Failed to determine home dir")
		return
	}

	uploadsDir := filepath.Join(path, utils.GetUploadsDirName())

	// List all files in the uploads directory
	files, err := os.ReadDir(uploadsDir)
	if err != nil {
		response.WithJSON(w, http.StatusInternalServerError, "Failed to list uploads")
		return
	}

	var fileNames []UploadsList
	for _, file := range files {
		fileInfo, err := file.Info()
		var size int64
		if err != nil {
			size = -1
		} else {
			size = fileInfo.Size()
		}

		fileNames = append(fileNames, UploadsList{
			Name: file.Name(),
			Size: size,
		})
	}

	response.WithJSON(w, http.StatusOK, fileNames)
}

type DeletePayload struct {
	File string `json:"file"`
}

func HandleDeleteUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.WithJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Decode the JSON payload
	var payload DeletePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.WithJSON(w, http.StatusBadRequest, "Failed to parse JSON payload")
		return
	}

	filename := payload.File

	if !utils.IsPathSafe(filename) {
		response.WithJSON(w, http.StatusBadRequest, "Path is not safe")
		return
	}

	path, err := utils.GetHomeDir()
	if err != nil {
		response.WithJSON(w, http.StatusInternalServerError, "Failed to determine home dir")
		return
	}

	uploadsDir := filepath.Join(path, utils.GetUploadsDirName())

	// Check if the file exists before attempting to delete
	filePath := filepath.Join(uploadsDir, filename)

	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist
			response.WithJSON(w, http.StatusNotFound, "File not found")
		} else {
			// Other error occurred while accessing the file
			response.WithJSON(w, http.StatusInternalServerError, "Failed to access the file")
		}
		return
	}

	// Delete the file
	err = os.Remove(filePath)
	if err != nil {
		response.WithJSON(w, http.StatusInternalServerError, "Failed to delete the file")
		return
	}

	response.WithJSON(w, http.StatusOK, nil)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WithJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Determine Path
	path, err := utils.GetHomeDir()
	if err != nil {
		response.WithJSON(w, http.StatusInternalServerError, "Failed to determine home dir")
		return
	}

	err = r.ParseMultipartForm(10 << 30) // 1GB maximum file size
	if err != nil {
		response.WithJSON(w, http.StatusInternalServerError, "Failed to parse form")
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		response.WithJSON(w, http.StatusBadRequest, "Failed to get file from form")
		return
	}
	defer file.Close()

	f, err := os.OpenFile(filepath.Join(path+"/"+utils.GetUploadsDirName(), handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		response.WithJSON(w, http.StatusInternalServerError, "Failed to save file")
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		response.WithJSON(w, http.StatusInternalServerError, "Failed to write file content")
		return
	}

	response.WithJSON(w, http.StatusOK, nil)
}
