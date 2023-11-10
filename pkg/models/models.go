package models

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" db:"username" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

type Group struct {
	ID          int    `json:"-" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Description string `json:"description" db:"description" binding:"required"`
}

type Role struct {
	ID   int    `json:"-" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}

type UserGroup struct {
	ID      int `json:"-" db:"id"`
	UserID  int `json:"userID" db:"userID" binding:"required"`
	GroupID int `json:"groupID" db:"groupID" binding:"required"`
	RoleID  int `json:"roleID" db:"roleID" binding:"required"`
}

type PlayList struct {
	ID      int    `json:"-" db:"id"`
	Name    string `json:"name" db:"name" binding:"required"`
	GroupID int    `json:"groupID" db:"groupID" binding:"required"`
}

type UpdateGroupInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type UpdateUserGroupInput struct {
	UserID  *int `json:"userID"`
	GroupID *int `json:"groupID"`
}

type UpdateRoleInput struct {
	Name *string `json:"name"`
}

type UpdatePlayListInput struct {
	Name    *string `json:"name"`
	GroupID *int    `json:"groupID"`
}