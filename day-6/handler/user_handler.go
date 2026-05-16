package handler

import (
	"fmt"
	"net/http"
	"strconv"

	svc "github.com/nathanaelsanilo/my-app/service"
)

type UserHandler interface {
	GetListUsersHandler(w http.ResponseWriter, r *http.Request)
	GetUserByIdHandler(w http.ResponseWriter, r *http.Request)
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
	UpdateUserHandler(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService svc.UserService
}

func NewUserHandler(svc svc.UserService) UserHandler {
	return &userHandler{userService: svc}
}

func (h *userHandler) GetListUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "log::GET list users::%s \n", h.userService.GetListUsers())
}

func (h *userHandler) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("user_id") // read path variable
	val, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Invalid user id!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "log::GET user id::%s \n", h.userService.GetUserById(val))
}

func (h *userHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "log::POST create user handler::%s\n", h.userService.CreateUser())
}

func (h *userHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("user_id") // read path variable
	if userId == "" {
		http.Error(w, "Invalid user id!", http.StatusBadRequest)
		return
	}

	fmt.Printf("log::UpdateUserHandler::userId::%s", userId)
	val, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Printf("error: %v", err)
		http.Error(w, "Invalid user id!", http.StatusBadRequest)
		return
	}
	payload := "mock"

	fmt.Fprintf(w, "log::PUT update user handler!::%s \n", h.userService.UpdateUserById(val, payload))
}
