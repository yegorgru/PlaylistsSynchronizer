package services

import (
	"PlaylistsSynchronizer/pkg/models"
	"PlaylistsSynchronizer/pkg/repositories"
)

type GroupService struct {
	repoGroup repositories.Group
	repoRole  repositories.Role
}

func NewGroupService(repoGroup repositories.Group, repoRole repositories.Role) *GroupService {
	return &GroupService{
		repoGroup: repoGroup,
		repoRole:  repoRole,
	}
}

func (s *GroupService) Create(userID int, group models.UserCreateGroup) (int, error) {
	role, err := s.repoRole.GetByName("ADMIN")
	if err != nil {
		return 0, err
	}
	return s.repoGroup.Create(userID, role.ID, group)
}

func (s *GroupService) GetAll() ([]models.Group, error) {
	return s.repoGroup.GetAll()
}

func (s *GroupService) GetById(id int) (models.Group, error) {
	return s.repoGroup.GetById(id)
}

func (s *GroupService) Update(id int, group models.UpdateGroupInput) error {
	if err := group.Validate(); err != nil {
		return err
	}
	return s.repoGroup.Update(id, group)
}

func (s *GroupService) Delete(id int) error {
	return s.repoGroup.Delete(id)
}
