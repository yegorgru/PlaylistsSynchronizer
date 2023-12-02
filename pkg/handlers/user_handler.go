package handlers

import (
	"PlaylistsSynchronizer/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	user, err := h.services.Authorization.GetUserByID(id)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if user == nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "there is no user with such id")
		return
	}
	c.JSON(http.StatusOK, user)
}
