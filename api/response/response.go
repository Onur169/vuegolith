package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Ack  string      `json:"ack"`
	Data interface{} `json:"data"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WithJSON(w http.ResponseWriter, status int, data any) {
	response := APIResponse{
		Ack:  "success",
		Data: data,
	}

	err := WriteJSON(w, status, response)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}
