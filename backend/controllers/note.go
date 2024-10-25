package controllers

import (
	"github.com/VitaliiVinnychenko/neurocast-task/backend/models"
	"github.com/VitaliiVinnychenko/neurocast-task/backend/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
)

// CreateNewNote godoc
// @Summary      Create Note
// @Description  creates a new note
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        req  body      models.NoteRequest true "Note Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /notes [post]
// @Security     ApiKeyAuth
func CreateNewNote(c *gin.Context) {
	var requestBody models.NoteRequest
	_ = c.ShouldBindBodyWith(&requestBody, binding.JSON)

	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	note, err := services.CreateNote(requestBody.Title, requestBody.Content)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusCreated
	response.Success = true
	response.Data = gin.H{"note": note}
	response.SendResponse(c)
}

// GetNotes godoc
// @Summary      Get Notes
// @Description  gets user notes with pagination
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        page  query    string  false  "Switch page by 'page'"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /notes [get]
// @Security     ApiKeyAuth
func GetNotes(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	pageQuery := c.DefaultQuery("page", "0")
	page, _ := strconv.Atoi(pageQuery)
	limit := 5

	notes, _ := services.GetNotes(page, limit)
	hasPrev := page > 0
	hasNext := len(notes) > limit

	if hasNext {
		notes = notes[:limit]
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{"notes": notes, "prev": hasPrev, "next": hasNext}
	response.SendResponse(c)
}

// GetOneNote godoc
// @Summary      Get a note
// @Description  get note by id
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Note ID"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /notes/{id} [get]
// @Security     ApiKeyAuth
func GetOneNote(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	idHex := c.Param("id")
	noteId, _ := primitive.ObjectIDFromHex(idHex)

	note, err := services.GetNoteById(noteId)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{"note": note}
	response.SendResponse(c)
}

// UpdateNote godoc
// @Summary      Update a note
// @Description  updates a note by id
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        id     path    string  true  "Note ID"
// @Param        req    body    models.NoteRequest true "Note Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /notes/{id} [put]
// @Security     ApiKeyAuth
func UpdateNote(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	idHex := c.Param("id")
	noteId, _ := primitive.ObjectIDFromHex(idHex)

	var noteRequest models.NoteRequest
	_ = c.ShouldBindBodyWith(&noteRequest, binding.JSON)

	err := services.UpdateNote(noteId, &noteRequest)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.SendResponse(c)
}

// DeleteNote godoc
// @Summary      Delete a note
// @Description  deletes note by id
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Note ID"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /notes/{id} [delete]
// @Security     ApiKeyAuth
func DeleteNote(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	idHex := c.Param("id")
	noteId, _ := primitive.ObjectIDFromHex(idHex)

	err := services.DeleteNote(noteId)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.SendResponse(c)
}
