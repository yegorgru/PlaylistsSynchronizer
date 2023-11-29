package services

import (
	"PlaylistsSynchronizer/pkg/api_services"
	"PlaylistsSynchronizer/pkg/models"
	"PlaylistsSynchronizer/pkg/repositories"
	"github.com/gin-gonic/gin"
)

type GroupService struct {
	reposAuth  repositories.Authorization
	repoGroup  repositories.Group
	repoRole   repositories.Role
	repoToken  repositories.Token
	apiService api_services.ApiService
}

func NewGroupService(reposAuth repositories.Authorization, repoGroup repositories.Group, repoRole repositories.Role,
	repoToken repositories.Token) *GroupService {
	return &GroupService{
		reposAuth:  reposAuth,
		repoGroup:  repoGroup,
		repoRole:   repoRole,
		repoToken:  repoToken,
		apiService: *api_services.NewApiService(repoToken),
	}
}

func (s *GroupService) Create(c *gin.Context, userID int, group models.UserCreateGroupInput) (int, error) {
	role, err := s.repoRole.GetByName("ADMIN")
	if err != nil {
		return 0, err
	}
	var playListID string
	playList := models.PlayList{Name: group.PlayListName, Description: group.PlayListDescription}
	switch group.Platform {
	case models.Spotify:
		spotify := s.apiService.GetSpotifyServiceApi()
		userSpotify, err := s.reposAuth.GetUserSpotifyByID(userID)
		if err != nil {
			return 0, err
		}
		spotifyData := models.SpotifyData{Token: userSpotify.AccessToken, SpotifyUri: userSpotify.SpotifyUri}
		playListID, err = spotify.CreatePlayList(spotifyData, playList)
		if err != nil {
			return 0, err
		}
	case models.YouTubeMusic:
		youTubeMusic := s.apiService.GetYouTubeMusicApiServiceApi()
		token, err := s.repoToken.GetYouTubeMusicToken(userID)
		if err != nil {
			return 0, err
		}
		playListID, err = youTubeMusic.CreatePlayList(token.AccessToken, playList)
		if err != nil {
			return 0, err
		}
	}
	group.PlayListID = playListID
	id, err := s.repoGroup.Create(userID, role.ID, group)
	if err != nil {
		return 0, err
	}
	return id, err
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
