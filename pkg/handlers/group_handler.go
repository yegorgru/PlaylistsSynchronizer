package handlers

import (
	"PlaylistsSynchronizer/pkg/models"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func (h *Handler) createGroup(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	platform, err := getUserPlatform(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, models.UserCreateGroupInput{}); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input models.UserCreateGroupInput
	if err := c.BindJSON(&input); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	input.Platform = platform
	id, err := h.services.Group.Create(userID, input)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllGroups(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	groups, err := h.services.Group.GetAll()
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, groups)
}

func (h *Handler) getGroupById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	group, err := h.services.Group.GetById(id)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, group)
}

func (h *Handler) updateGroup(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid groupID param")
		return
	}
	isGroupExist, err := h.services.Group.GetById(groupID)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if isGroupExist == nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "there is no group with such id")
		return
	}
	isValidUser, err := h.isValidAdmin(groupID, userID)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "user validation error")
		return
	}
	if isValidUser {
		body, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		if err := h.validateJSONTags(body, models.UpdateGroupInput{}); err != nil {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		var input models.UpdateGroupInput
		if err := c.BindJSON(&input); err != nil {
			models.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
			return
		}

		if err := h.services.Group.Update(groupID, input); err != nil {
			models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, models.StatusResponse{Status: "ok"})
	} else {
		models.NewErrorResponse(c, http.StatusForbidden, "invalid permission")
		return
	}
}

func (h *Handler) deleteGroup(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid groupID param")
		return
	}
	isGroupExist, err := h.services.Group.GetById(groupID)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if isGroupExist == nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "there is no group with such id")
		return
	}
	isValidUser, err := h.isValidUserRole(groupID, userID, "SUPER ADMIN")
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "user validation error")
		return
	}
	if isValidUser {
		if err := h.services.Group.Delete(groupID); err != nil {
			models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, models.StatusResponse{Status: "ok"})
	} else {
		models.NewErrorResponse(c, http.StatusForbidden, "invalid permission")
		return
	}
}
