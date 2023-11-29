package handlers

import (
	"PlaylistsSynchronizer/pkg/models"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (h *Handler) addTrack(c *gin.Context) {
	var input models.AddTrackInput

	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, models.AddTrackInput{}); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.BindJSON(&input); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Track.Add(input)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
