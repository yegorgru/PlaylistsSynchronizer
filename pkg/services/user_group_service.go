package services

import (
	"PlaylistsSynchronizer/pkg/models"
	"PlaylistsSynchronizer/pkg/repositories"
)

type UserGroupService struct {
	repoUserGroup repositories.UserGroup
	repoRole      repositories.Role
}

func NewUserGroupService(repoUserGroup repositories.UserGroup, repoRole repositories.Role) *UserGroupService {
	return &UserGroupService{
		repoUserGroup: repoUserGroup,
		repoRole:      repoRole,
	}
}

func (s *UserGroupService) Create(group models.UserGroup) (int, error) {
	role, err := s.repoRole.GetByName("USER")
	group.RoleID = role.ID
	if err != nil {
		return 0, err
	}
	return s.repoUserGroup.Create(group)
}

func (s *UserGroupService) GetAll() ([]models.UserGroup, error) {
	return s.repoUserGroup.GetAll()
}

func (s *UserGroupService) GetById(id int) (models.UserGroup, error) {
	return s.repoUserGroup.GetById(id)
}

func (s *UserGroupService) Update(id int, group models.UpdateUserGroupInput) error {
	if err := group.Validate(); err != nil {
		return err
	}
	return s.repoUserGroup.Update(id, group)
}

func (s *UserGroupService) Delete(id int) error {
	return s.repoUserGroup.Delete(id)
}
