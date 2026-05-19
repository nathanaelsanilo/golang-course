package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nathanaelsanilo/my-app/dto"
	"github.com/nathanaelsanilo/my-app/helper"
	"github.com/nathanaelsanilo/my-app/service"
)

type NoteHandler struct {
	svc service.NoteService
}

func NewNoteHandler(svc service.NoteService) *NoteHandler {
	return &NoteHandler{svc: svc}
}

func (s *NoteHandler) GetList(w http.ResponseWriter, r *http.Request) {
	helper.WriteSuccess(w, s.svc.GetList())
}

func (s *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	// get request body
	var body dto.NoteCreateDto
	helper.GetBody(r, &body)
	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&body)
	// if err != nil {
	// 	helper.WriteError(w, "Type miss match!")
	// 	return
	// }

	if body.Validate() {
		helper.WriteError(w, "Missing required fields value!")
		return
	}

	helper.WriteSuccess(w, s.svc.Create())
}

func (s *NoteHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intId, _ := strconv.Atoi(id)
	if intId%2 == 0 {
		helper.WriteSuccess(w, s.svc.GetDetail(chi.URLParam(r, "id")))
		return
	}

	helper.WriteError(w, "ID Not found!")
}

func (s *NoteHandler) Update(w http.ResponseWriter, r *http.Request) {
	helper.WriteSuccess(w, s.svc.Update(chi.URLParam(r, "id")))
}

func (s *NoteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	helper.WriteError(w, "Something went wrong!")
}
