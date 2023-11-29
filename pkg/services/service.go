package services

import (
	"PlaylistsSynchronizer/pkg/models"
	"PlaylistsSynchronizer/pkg/repositories"
	"github.com/gin-gonic/gin"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	CreateUserSpotify(spotifyUri string, token models.ApiToken, user models.User) (int, error)
	CreateUserYouTubeMusic(token string, user models.User) (int, error)
	GetUser(email string) (*models.User, error)
	GenerateToken(email string) (string, error)
	ParseToken(token string) (*models.UserClaims, error)
	RefreshToken(userId int) (string, error)
}

type Group interface {
	Create(c *gin.Context, userID int, group models.UserCreateGroupInput) (int, error)
	GetAll() ([]models.Group, error)
	GetById(id int) (models.Group, error)
	Update(id int, group models.UpdateGroupInput) error
	Delete(id int) error
}

type UserGroup interface {
	Create(c *gin.Context, platform string, userGroup models.UserGroup) (int, error)
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

type Track interface {
	Add(input models.AddTrackInput) (int, error)
	DeleteFromPlayList(trackID string) error
}

type Token interface {
	Create(token models.Token) (int, error)
	GetByToken(token string) (*models.Token, error)
	Update(token string) error
	UpdateYouTubeAccessToken(token string, userID int) error
	RevokeAllUserTokens(userID int) error
}

type Service struct {
	Authorization
	Group
	UserGroup
	Role
	PlayList
	Track
	Token
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Group:         NewGroupService(repos.Authorization, repos.Group, repos.Role, repos.Token),
		UserGroup: NewUserGroupService(repos.Authorization, repos.UserGroup, repos.PlayList, repos.Role, repos.Track,
			repos.Token),
		Role:     NewRoleService(repos.Role),
		PlayList: NewPlayListService(repos.PlayList),
		Track:    NewTrackService(repos.Track, repos.UserGroup, repos.Token, repos.PlayList),
		Token:    NewTokenService(repos.Token),
	}
}
