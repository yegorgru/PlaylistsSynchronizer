package services

import (
	"PlaylistsSynchronizer/pkg/models"
	"PlaylistsSynchronizer/pkg/repositories"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Group interface {
	Create(userID int, group models.UserCreateGroup) (int, error)
	GetAll() ([]models.Group, error)
	GetById(id int) (models.Group, error)
	Update(id int, group models.UpdateGroupInput) error
	Delete(id int) error
}

type UserGroup interface {
	Create(userGroup models.UserGroup) (int, error)
	GetAll() ([]models.UserGroup, error)
	GetById(id int) (models.UserGroup, error)
	Update(id int, group models.UpdateUserGroupInput) error
	Delete(id int) error
}

type Role interface {
	Create(role models.Role) (int, error)
	GetAll() ([]models.Role, error)
	GetById(id int) (models.Role, error)
	Update(id int, role models.UpdateRoleInput) error
	Delete(id int) error
}

type PlayList interface {
	Create(playList models.PlayList) (int, error)
	GetAll() ([]models.PlayList, error)
	GetById(id int) (models.PlayList, error)
	Update(id int, group models.UpdatePlayListInput) error
	Delete(id int) error
}

type Service struct {
	Authorization
	Group
	UserGroup
	Role
	PlayList
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Group:         NewGroupService(repos.Group, repos.Role),
		UserGroup:     NewUserGroupService(repos.UserGroup, repos.Role),
		Role:          NewRoleService(repos.Role),
		PlayList:      NewPlayListService(repos.PlayList),
	}
}
