package services

import (
	"PlaylistsSynchronizer/pkg/models"
	"PlaylistsSynchronizer/pkg/repositories"
)

type PlayListService struct {
	repo repositories.PlayList
}

func NewPlayListService(repo repositories.PlayList) *PlayListService {
	return &PlayListService{
		repo: repo,
	}
}

func (s *PlayListService) Create(playList models.PlayList) (int, error) {
	return s.repo.Create(playList)
}

func (s *PlayListService) GetAll() ([]models.PlayList, error) {
	return s.repo.GetAll()
}

func (s *PlayListService) GetById(id int) (models.PlayList, error) {
	return s.repo.GetById(id)
}

func (s *PlayListService) Update(id int, role models.UpdatePlayListInput) error {
	if err := role.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, role)
}

func (s *PlayListService) Delete(id int) error {
	return s.repo.Delete(id)
}
