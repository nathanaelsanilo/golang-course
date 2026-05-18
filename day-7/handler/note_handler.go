package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nathanaelsanilo/my-app/service"
)

type NoteHandler struct {
	svc service.NoteService
}

func NewNoteHandler(svc service.NoteService) *NoteHandler {
	return &NoteHandler{svc: svc}
}

func (s *NoteHandler) GetList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", s.svc.GetList())
}

func (s *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", s.svc.Create())
}

func (s *NoteHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", s.svc.GetDetail(chi.URLParam(r, "id")))
}

func (s *NoteHandler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", s.svc.Update(chi.URLParam(r, "id")))
}

func (s *NoteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", s.svc.Delete(chi.URLParam(r, "id")))
}
