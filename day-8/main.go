package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nathanaelsanilo/my-app/dto"
	"github.com/nathanaelsanilo/my-app/handler"
	"github.com/nathanaelsanilo/my-app/helper"
	"github.com/nathanaelsanilo/my-app/service"
)

func NewUserDto(id string) dto.UserDto {
	intId, _ := strconv.Atoi(id)
	return dto.UserDto{ID: intId, Name: "Nathan", Email: "", Password: "0000", DateOfBirth: time.Now()}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	userId := chi.URLParam(r, "id")
	userDto := NewUserDto(userId)
	// use json.Marshal if you need to modify it
	// since the return is store in memory
	userDtoBuffer, err := json.Marshal(userDto)
	if err != nil {
		http.Error(w, fmt.Errorf("Serialization failed! %w", err).Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User dto: %v", userDtoBuffer)
}

func getUserDetail(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	userDto := NewUserDto(userId)
	// fmt.Fprintf(w, "User : %v", userDto)
	w.Header().Set("Content-type", "application/json")

	// streams to io.Writer directly
	json.NewEncoder(w).Encode(userDto)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UserCreateDto

	// limit request body size
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Printf("req.email: %s\n", req.Email)
	fmt.Printf("req.name: %s\n", req.Name)
	json.NewEncoder(w).Encode(req)
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	isEven := rand.IntN(100)%2 == 0

	if isEven {
		helper.WriteSuccess(w, "Get all notes!")
		return
	}

	helper.WriteError(w, "Service unavailable!")
}

func main() {
	r := chi.NewRouter()

	r.Get("/", getUsers)
	r.Get("/{id}", getUserDetail)
	r.Post("/", createUser)

	noteService := service.NewNoteService()
	noteHandler := handler.NewNoteHandler(noteService)

	r.Route("/notes", func(r chi.Router) {
		r.Get("/", noteHandler.GetList)
		r.Get("/{id}", noteHandler.GetDetail)
		r.Post("/", noteHandler.Create)
		r.Put("/{id}", noteHandler.Update)
		r.Delete("/{id}", noteHandler.Delete)
	})

	http.ListenAndServe(":8080", r)
}
