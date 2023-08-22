package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Ack  string `json:"ack"`
	Data any    `json:"data"`
}

type ACK string

const (
	ResponseSuccess ACK = "success"
	ResponseFail    ACK = "fail"
)

func WriteJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}

func WithJSON(w http.ResponseWriter, ack ACK, statusCode int, data any) {
	response := APIResponse{
		Ack:  string(ack),
		Data: data,
	}

	err := WriteJSON(w, statusCode, response)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}
