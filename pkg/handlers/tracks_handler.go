package handlers

import (
	"PlaylistsSynchronizer/pkg/models"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func (h *Handler) addTrack(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}
	playListID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid playListID param")
		return
	}
	var input models.AddTrackInput

	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	if err := h.validateJSONTags(body, models.AddTrackInput{}); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := c.BindJSON(&input); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	playList, err := h.services.PlayList.GetById(playListID)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if playList == nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "there is no playlist with such id")
		return
	}
	isValidUser, err := h.isValidUser(playList.GroupID, userID)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "user validation error")
		return
	}
	if isValidUser {
		input.GroupID = playList.GroupID
		input.PlayListID = playList.ID
		input.UserID = userID
		id, err := h.services.Track.Add(input)
		if err != nil {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	} else {
		models.NewErrorResponse(c, http.StatusBadRequest, "you are not member of the group")
		return
	}
}

func (h *Handler) deleteTrack(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}
	playListID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid playListID param")
		return
	}
	trackID, err := strconv.Atoi(c.Param("trackID"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "invalid trackID param")
		return
	}
	isPlayListExist, err := h.services.PlayList.GetById(playListID)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if isPlayListExist == nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "there is no playlist with such id")
		return
	}
	isTrackExist, err := h.services.Track.GetByPlayListTrackID(playListID, trackID)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if isTrackExist == nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "there is no track with such id")
		return
	}
	isValidUser, err := h.isValidAdmin(isPlayListExist.GroupID, userID)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, "user validation error")
		return
	}
	if isValidUser {
		err := h.services.Track.DeleteFromPlayList(isPlayListExist.GroupID, playListID, trackID)
		if err != nil {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, models.StatusResponse{Status: "ok"})
	} else {
		models.NewErrorResponse(c, http.StatusForbidden, "invalid permission")
		return
	}
}
