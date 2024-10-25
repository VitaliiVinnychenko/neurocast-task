package models

import validation "github.com/go-ozzo/ozzo-validation"

type NoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (a NoteRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required),
		validation.Field(&a.Content, validation.Required),
	)
}
