package services

import (
	"PlaylistsSynchronizer/pkg/models"
	"PlaylistsSynchronizer/pkg/repositories"
)

type UserGroupService struct {
	repo repositories.UserGroup
}

func NewUserGroupService(repo repositories.UserGroup) *UserGroupService {
	return &UserGroupService{
		repo: repo,
	}
}

func (s *UserGroupService) Create(group models.UserGroup) (int, error) {
	return s.repo.Create(group)
}

func (s *UserGroupService) GetAll() ([]models.UserGroup, error) {
	return s.repo.GetAll()
}

func (s *UserGroupService) GetById(id int) (models.UserGroup, error) {
	return s.repo.GetById(id)
}

func (s *UserGroupService) Update(id int, group models.UpdateUserGroupInput) error {
	if err := group.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, group)
}

func (s *UserGroupService) Delete(id int) error {
	return s.repo.Delete(id)
}
