package models

import (
	"errors"
	"io"
)

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
	Platform string `json:"platform" db:"platform" binding:"required"`
}

type Group struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Description string `json:"description" db:"description" binding:"required"`
}

type Role struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}

type UserSpotify struct {
	ID           int    `json:"id" db:"id"`
	UserID       int    `json:"userID" db:"userid"`
	SpotifyUri   string `json:"spotifyUri" db:"spotifyuri"`
	AccessToken  string `json:"accessToken" db:"accesstoken"`
	RefreshToken string `json:"refreshToken" db:"refreshtoken"`
}

type UserGroup struct {
	ID         int    `json:"id" db:"id"`
	UserID     int    `json:"userID" db:"userid"`
	GroupID    int    `json:"groupID" db:"groupid"`
	RoleID     int    `json:"roleID" db:"roleid"`
	PlayListID string `json:"playListID" db:"playlistid"`
}

type UserGroupInput struct {
	GroupID int `json:"groupID" binding:"required"`
}

type PlayList struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Description string `json:"description" db:"description" binding:"required"`
	GroupID     int    `json:"groupID" db:"groupid" binding:"required"`
}

type SpotifyUris struct {
	Uris     []string `json:"uris"`
	Position int      `json:"position"`
}

type Track struct {
	ID             int    `json:"id" db:"id"`
	SpotifyUri     string `json:"spotifyUri" db:"spotifyuri"`
	YouTubeMusicID string `json:"youTubeMusicID" db:"youtubemusicid"`
}

type AddTrackInput struct {
	SpotifyUri     string `json:"spotifyUri" binding:"required"`
	YouTubeMusicID string `json:"youTubeMusicID" binding:"required"`
	GroupID        int    `json:"groupID" binding:"required"`
	PlayListID     int
}

type UserCreateGroupInput struct {
	ID                  int    `json:"id"`
	GroupName           string `json:"groupName" binding:"required"`
	PlayListName        string `json:"playListName" binding:"required"`
	GroupDescription    string `json:"groupDescription" binding:"required"`
	PlayListDescription string `json:"playListDescription" binding:"required"`
	PlayListID          string
	Platform            string
}

type UserGroupToken struct {
	ID         int    `json:"id" db:"id"`
	UserID     int    `json:"userID" db:"userid"`
	SpotifyUri string `json:"spotifyUri" db:"spotifyuri"`
	GroupID    int    `json:"groupID" db:"groupid"`
	RoleID     int    `json:"roleID" db:"roleid"`
	Token      string `json:"token" db:"accesstoken"`
	PlayListID string `json:"playListID" db:"playlistid"`
}

type UserClaims struct {
	UserID       int    `json:"userID"`
	UserPlatform string `json:"userPlatform"`
}

type UpdateGroupInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type UpdateUserGroupInput struct {
	UserID     *int    `json:"userID"`
	GroupID    *int    `json:"groupID"`
	RoleID     *int    `json:"roleID"`
	PlayListID *string `json:"playListIDs"`
}

type UpdateRoleInput struct {
	Name *string `json:"name"`
}

type UpdatePlayListInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	GroupID     *int    `json:"groupID"`
}

type ApiToken struct {
	AccessToken  string `json:"accessToken" db:"accesstoken"`
	RefreshToken string `json:"refreshToken" db:"refreshtoken"`
}

type Token struct {
	ID         int    `json:"id" db:"id"`
	TokenValue string `json:"tokenValue" db:"tokenvalue" binding:"required"`
	Revoked    bool   `json:"revoked" db:"revoked"`
	UserID     int    `json:"userID" db:"userid"`
}

type RefreshToken struct {
	Query string
	Body  io.Reader
}

type RefreshTokenInput struct {
	UserId int `json:"userId" binding:"required"`
}

type SpotifyData struct {
	Token      string
	SpotifyUri string
}

type ApiTrackID struct {
	SpotifyUri   string
	YouTubeMusic string
}

func (i UpdateGroupInput) Validate() error {
	if i.Name == nil && i.Description == nil {
		return errors.New("update structure has no value")
	}
	return nil
}

func (i UpdateUserGroupInput) Validate() error {
	if i.UserID == nil && i.GroupID == nil && i.RoleID == nil && i.PlayListID == nil {
		return errors.New("update structure has no value")
	}
	return nil
}

func (i UpdateRoleInput) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no value")
	}
	return nil
}

func (i UpdatePlayListInput) Validate() error {
	if i.Name == nil && i.Description == nil && i.GroupID == nil {
		return errors.New("update structure has no value")
	}
	return nil
}
