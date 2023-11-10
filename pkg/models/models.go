package models

import "errors"

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" db:"username" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
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

type UserGroup struct {
	ID      int `json:"id" db:"id"`
	UserID  int `json:"userID" db:"userid"`
	GroupID int `json:"groupID" db:"groupid"`
	RoleID  int `json:"roleID" db:"roleid"`
}

type PlayList struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name" binding:"required"`
	GroupID int    `json:"groupID" db:"groupid" binding:"required"`
}

type UserCreateGroup struct {
	ID           int    `json:"-"`
	GroupName    string `json:"groupName"`
	PlayListName string `json:"playListName"`
	Description  string `json:"description"`
}

type UpdateGroupInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type UpdateUserGroupInput struct {
	UserID  *int `json:"userID"`
	GroupID *int `json:"groupID"`
	RoleID  *int `json:"roleID"`
}

type UpdateRoleInput struct {
	Name *string `json:"name"`
}

type UpdatePlayListInput struct {
	Name    *string `json:"name"`
	GroupID *int    `json:"groupID"`
}

func (i UpdateGroupInput) Validate() error {
	if i.Name == nil && i.Description == nil {
		return errors.New("update structure has no value")
	}
	return nil
}

func (i UpdateUserGroupInput) Validate() error {
	if i.UserID == nil && i.GroupID == nil {
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
	if i.Name == nil && i.GroupID == nil {
		return errors.New("update structure has no value")
	}
	return nil
}
