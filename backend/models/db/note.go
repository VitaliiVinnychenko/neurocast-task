package models

import (
	"github.com/kamva/mgm/v3"
)

type Note struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string             `json:"title" bson:"title"`
	Content          string             `json:"content" bson:"content"`
}

func NewNote(title string, content string) *Note {
	return &Note{
		Title:   title,
		Content: content,
	}
}

func (model *Note) CollectionName() string {
	return "notes"
}
