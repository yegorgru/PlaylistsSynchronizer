package handlers

import (
	"PlaylistsSynchronizer/pkg/models"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func (h *Handler) createRole(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	var input models.Role
	if err := c.BindJSON(&input); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Role.Create(input)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllRoleResponse struct {
	Data []models.Role `json:data`
}

func (h *Handler) getAllRole(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	roles, err := h.services.Role.GetAll()
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllRoleResponse{
		Data: roles,
	})
}

func (h *Handler) getRoleById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	role, err := h.services.Role.GetById(id)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, role)
}

func (h *Handler) updateRole(c *gin.Context) {
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
	if err := h.validateJSONTags(body, models.UpdateRoleInput{}); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input models.UpdateRoleInput

	if err := c.BindJSON(&input); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if err := h.services.Role.Update(id, input); err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, models.StatusResponse{Status: "ok"})
}

func (h *Handler) deleteRole(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Role.Delete(id); err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, models.StatusResponse{Status: "ok"})
}
