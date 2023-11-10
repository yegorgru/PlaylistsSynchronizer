package repositories

import (
	"PlaylistsSynchronizer/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Group interface {
	Create(userID, roleID int, group models.Group) (int, error)
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
	Create(group models.Role) (int, error)
	GetAll() ([]models.Role, error)
	GetById(id int) (models.Role, error)
	Update(id int, group models.UpdateRoleInput) error
	Delete(id int) error
}

type PlayList interface {
	Create(group models.PlayList) (int, error)
	GetAll() ([]models.PlayList, error)
	GetById(id int) (models.PlayList, error)
	Update(id int, group models.UpdatePlayListInput) error
	Delete(id int) error
}

type Repository struct {
	Authorization
	Group
	UserGroup
	Role
	PlayList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Group:         NewGroupPostgres(db),
		UserGroup:     NewUserGroupPostgres(db),
		Role:          NewRolePostgres(db),
		PlayList:      NewPlayListPostgres(db),
	}
}
