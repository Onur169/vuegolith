package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Ack  string      `json:"ack"`
	Data interface{} `json:"data"`
}

func WithJSON(w http.ResponseWriter, status int, data interface{}) {
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
