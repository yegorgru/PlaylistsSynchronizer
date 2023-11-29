package handlers

import (
	"PlaylistsSynchronizer/pkg/models"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func (h *Handler) createPlayList(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, models.PlayList{}); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input models.PlayList
	if err := c.BindJSON(&input); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.PlayList.Create(input)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllPlayList(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	playLists, err := h.services.PlayList.GetAll()
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, playLists)
}

func (h *Handler) getPlayListById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	playlist, err := h.services.PlayList.GetById(id)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, playlist)
}

func (h *Handler) updatePlayList(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, models.UpdatePlayListInput{}); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input models.UpdatePlayListInput

	if err := c.BindJSON(&input); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if err := h.services.PlayList.Update(id, input); err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, models.StatusResponse{Status: "ok"})
}

func (h *Handler) deletePlayList(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.PlayList.Delete(id); err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, models.StatusResponse{Status: "ok"})
}
