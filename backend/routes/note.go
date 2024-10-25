package routes

import (
	"github.com/VitaliiVinnychenko/neurocast-task/backend/controllers"
	"github.com/VitaliiVinnychenko/neurocast-task/backend/middlewares/validators"
	"github.com/gin-gonic/gin"
)

func NoteRoute(router *gin.RouterGroup) {
	notes := router.Group("/notes")
	{
		notes.POST(
			"",
			validators.CreateNoteValidator(),
			controllers.CreateNewNote,
		)

		notes.GET(
			"",
			validators.GetNotesValidator(),
			controllers.GetNotes,
		)

		notes.GET(
			"/:id",
			validators.PathIdValidator(),
			controllers.GetOneNote,
		)

		notes.PUT(
			"/:id",
			validators.PathIdValidator(),
			validators.UpdateNoteValidator(),
			controllers.UpdateNote,
		)

		notes.DELETE(
			"/:id",
			validators.PathIdValidator(),
			controllers.DeleteNote,
		)
	}
}
