package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	svc UserService
}

func NewUserHandler(svc UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (s *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body UserCreateReqDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := s.svc.Create(body)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteSuccess(w, http.StatusCreated, res)
}

func (s *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		WriteError(w, http.StatusBadRequest, "ID is invalid!")
		return
	}

	res, err := s.svc.GetById(userId)
	if err != nil {
		WriteError(w, http.StatusNotFound, "User not found!")
		return
	}

	WriteSuccess(w, http.StatusOK, res)
}
