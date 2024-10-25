package services

import (
	"errors"
	"github.com/VitaliiVinnychenko/neurocast-task/backend/models"
	db "github.com/VitaliiVinnychenko/neurocast-task/backend/models/db"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateNote create new note record
func CreateNote(title string, content string) (*db.Note, error) {
	note := db.NewNote(title, content)
	err := mgm.Coll(note).Create(note)
	if err != nil {
		return nil, errors.New("cannot create new note")
	}

	return note, nil
}

// GetNotes get notes list
func GetNotes() ([]db.Note, error) {
	var notes []db.Note

	err := mgm.Coll(&db.Note{}).SimpleFind(
		&notes,
		bson.M{},
	)

	if err != nil {
		return nil, errors.New("cannot find notes")
	}

	return notes, nil
}

func GetNoteById(noteId primitive.ObjectID) (*db.Note, error) {
	note := &db.Note{}
	err := mgm.Coll(note).First(bson.M{field.ID: noteId}, note)
	if err != nil {
		return nil, errors.New("cannot find note")
	}

	return note, nil
}

// UpdateNote updates a note with id
func UpdateNote(noteId primitive.ObjectID, request *models.NoteRequest) error {
	note := &db.Note{}
	err := mgm.Coll(note).FindByID(noteId, note)
	if err != nil {
		return errors.New("cannot find note")
	}

	note.Title = request.Title
	note.Content = request.Content
	err = mgm.Coll(note).Update(note)

	if err != nil {
		return errors.New("cannot update")
	}

	return nil
}

// DeleteNote delete a note with id
func DeleteNote(noteId primitive.ObjectID) error {
	deleteResult, err := mgm.Coll(&db.Note{}).DeleteOne(mgm.Ctx(), bson.M{field.ID: noteId})

	if err != nil || deleteResult.DeletedCount <= 0 {
		return errors.New("cannot delete note")
	}

	return nil
}
