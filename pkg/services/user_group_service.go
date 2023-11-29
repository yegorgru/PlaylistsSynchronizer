package services

import (
	"PlaylistsSynchronizer/pkg/api_services"
	"PlaylistsSynchronizer/pkg/models"
	"PlaylistsSynchronizer/pkg/repositories"
	"github.com/gin-gonic/gin"
)

type UserGroupService struct {
	repoAuth      repositories.Authorization
	repoUserGroup repositories.UserGroup
	repoPlayList  repositories.PlayList
	repoRole      repositories.Role
	repoTrack     repositories.Track
	repoToken     repositories.Token
	apiService    api_services.ApiService
}

func NewUserGroupService(repoAuth repositories.Authorization, repoUserGroup repositories.UserGroup,
	repoPlayList repositories.PlayList, repoRole repositories.Role, repoTrack repositories.Track,
	repoToken repositories.Token) *UserGroupService {
	return &UserGroupService{
		repoAuth:      repoAuth,
		repoUserGroup: repoUserGroup,
		repoPlayList:  repoPlayList,
		repoRole:      repoRole,
		repoTrack:     repoTrack,
		repoToken:     repoToken,
		apiService:    *api_services.NewApiService(repoToken),
	}
}

func (s *UserGroupService) Create(c *gin.Context, platform string, group models.UserGroup) (int, error) {
	role, err := s.repoRole.GetByName("USER")
	group.RoleID = role.ID
	if err != nil {
		return 0, err
	}
	playList, err := s.repoPlayList.GetByGroupId(group.GroupID)
	if err != nil {
		return 0, err
	}
	tracks, err := s.repoTrack.GetByPlayListID(playList.ID)
	if err != nil {
		return 0, err
	}
	var playListID string
	switch platform {
	case models.Spotify:
		spotify := s.apiService.GetSpotifyServiceApi()
		userSpotify, err := s.repoAuth.GetUserSpotifyByID(group.UserID)
		if err != nil {
			return 0, err
		}
		token, err := s.repoToken.GetSpotifyToken(userSpotify.SpotifyUri)
		spotifyData := models.SpotifyData{Token: token.AccessToken, SpotifyUri: userSpotify.SpotifyUri}
		playListID, err = spotify.CreatePlayList(spotifyData, playList)
		if err != nil {
			return 0, err
		}
		err = spotify.AddTrack(spotifyData, playListID, tracks)
		if err != nil {
			return 0, err
		}
	case models.YouTubeMusic:
		youTubeMusic := s.apiService.GetYouTubeMusicApiServiceApi()
		token, err := s.repoToken.GetYouTubeMusicToken(group.UserID)
		playListID, err = youTubeMusic.CreatePlayList(token.AccessToken, playList)
		if err != nil {
			return 0, err
		}
		err = youTubeMusic.AddTrack(token.AccessToken, playListID, tracks)
		if err != nil {
			return 0, err
		}
	}
	group.PlayListID = playListID
	return s.repoUserGroup.Create(group)
}

func (s *UserGroupService) GetAll() ([]models.UserGroup, error) {
	return s.repoUserGroup.GetAll()
}

func (s *UserGroupService) GetById(id int) (models.UserGroup, error) {
	return s.repoUserGroup.GetById(id)
}

func (s *UserGroupService) GetByGroupId(id int) ([]models.UserGroup, error) {
	return s.repoUserGroup.GetByGroupId(id)
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
