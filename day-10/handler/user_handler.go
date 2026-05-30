package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nathanaelsanilo/my-app/app"
	"github.com/nathanaelsanilo/my-app/models"
	"github.com/nathanaelsanilo/my-app/services"
	"github.com/nathanaelsanilo/my-app/utils"
)

type UserHandler struct {
	svc services.UserService
}

func NewUserHandler(svc services.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (s *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res := s.svc.GetAll()

	utils.WriteSuccess(w, http.StatusOK, res)
}

func (s *UserHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "user_id"))

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, errUserDetail := s.svc.GetDetail(userId)

	if errUserDetail != nil {
		if errors.Is(errUserDetail, app.ErrUserNotFound) {
			utils.WriteError(w, http.StatusNotFound, errUserDetail.Error())
			return
		}

		utils.WriteError(w, http.StatusInternalServerError, errUserDetail.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, user)
}

func (s *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto models.UserReqDto
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	res := s.svc.Create(dto)

	utils.WriteSuccess(w, http.StatusCreated, res)
}

func (s *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	userId, errUrlParam := strconv.Atoi(chi.URLParam(r, "user_id"))

	if errUrlParam != nil {
		utils.WriteError(w, http.StatusBadRequest, errUrlParam.Error())
		return
	}

	var dto models.UserReqDto
	errBody := json.NewDecoder(r.Body).Decode(&dto)

	if errBody != nil {
		utils.WriteError(w, http.StatusBadRequest, errBody.Error())
		return
	}

	res, errUpdate := s.svc.Update(userId, dto)

	if errUpdate != nil {
		if errors.Is(errUpdate, app.ErrUserNotFound) {
			utils.WriteError(w, http.StatusNotFound, errUpdate.Error())
			return
		}

		utils.WriteError(w, http.StatusInternalServerError, errUpdate.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, res)
}

func (s *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userId, errUrlParam := strconv.Atoi(chi.URLParam(r, "user_id"))

	if errUrlParam != nil {
		utils.WriteError(w, http.StatusBadRequest, errUrlParam.Error())
		return
	}

	err := s.svc.DeleteById(userId)

	if err != nil {

		if errors.Is(err, app.ErrUserNotFound) {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}

		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
