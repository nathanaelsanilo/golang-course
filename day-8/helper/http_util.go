package helper

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

// utility helper to write json response
// write directly to stream
func WriteJSON(w http.ResponseWriter, httpStatus int, v any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, msg string) {
	WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: msg})
}

func WriteSuccess(w http.ResponseWriter, v any) {
	WriteJSON(w, http.StatusOK, SuccessResponse{Data: v})
}

func GetBody(r *http.Request, out any) error {
	// var body any
	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&body)

	// if err != nil {
	// 	return err
	// }

	// return decoder.Decode(body)
	return json.NewDecoder(r.Body).Decode(out)
}
