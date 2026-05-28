package main

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

func WriteJson(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	WriteJson(w, status, ErrorResponse{Error: msg})
}

func WriteSuccess(w http.ResponseWriter, status int, v any) {
	WriteJson(w, status, SuccessResponse{Data: v})
}
