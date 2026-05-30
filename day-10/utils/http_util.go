package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Data any `json:"data"`
}

func WriteJson(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, statusCode int, msg string) {
	WriteJson(w, statusCode, ErrorResponse{Error: msg})
}

func WriteSuccess(w http.ResponseWriter, statusCode int, v any) {
	WriteJson(w, statusCode, SuccessResponse{Data: v})
}
