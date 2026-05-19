package dto

type NoteCreateDto struct {
	Content string `json:"content"`
}

func NewNoteCreateDto() NoteCreateDto {
	return NoteCreateDto{}
}

func (d *NoteCreateDto) Validate() bool {
	return d.Content == ""
}
