package service

type NoteService interface {
	Create() string
	GetList() string
	GetDetail(id string) string
	Update(id string) string
	Delete(id string) string
}

type noteService struct {
}

func NewNoteService() NoteService {
	return &noteService{}
}

func (s *noteService) Create() string {
	return "create a note!"
}

func (s *noteService) GetList() string {
	return "get list!"
}

func (s *noteService) GetDetail(id string) string {
	return "find a note! -> " + id
}

func (s *noteService) Update(id string) string {
	return "update a note! -> " + id
}

func (s *noteService) Delete(id string) string {
	return "delete a note! -> " + id
}
