package handlers

import (
	_ "PlaylistsSynchronizer/docs"
	"PlaylistsSynchronizer/pkg/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/logout", h.logout)
		auth.GET("/spotify-login", h.spotifyLogin)
		auth.GET("/spotify-callback", h.spotifyCallBack)
		auth.GET("/youtube-music-login", h.youTubeMusicLogin)
		auth.GET("/youtube-music-callback", h.youTubeMusicCallBack)
		auth.GET("/apple-music-login", h.appleMusicLogin)
		auth.GET("/apple-music-callback", h.appleMusicCallBack)
	}

	api := router.Group("/api", h.userIdentity)
	{
		groups := api.Group("/groups")
		{
			groups.POST("/", h.createGroup)
			groups.GET("/", h.getAllGroups)
			groups.GET("/:id", h.getGroupById)
			groups.PUT("/:id", h.updateGroup)
			groups.DELETE("/:id", h.deleteGroup)
			groups.POST("/:id/leave", h.leaveGroup)
			groups.POST("/:id/users", h.createUserGroup)
			groups.GET("/:id/users", h.getAllUserGroups)
			groups.GET("/:id/users/:userID", h.getUserGroupByUserId)
			groups.PUT("/:id/users/:userID", h.updateUserGroup)
			groups.DELETE("/:id/users/:userID", h.deleteUserGroup)
		}
		users := api.Group("/users")
		{
			users.GET("/:id", h.getUserByID)
		}
		roles := api.Group("/roles")
		{
			roles.POST("/", h.createRole)
			roles.GET("/", h.getAllRole)
			roles.GET("/:id", h.getRoleById)
			roles.PUT("/:id", h.updateRole)
			roles.DELETE("/:id", h.deleteRole)
		}
		playLists := api.Group("/playlists")
		{
			playLists.GET("/", h.getAllPlayList)
			playLists.GET("/:id", h.getPlayListById)
			playLists.PUT("/:id", h.updatePlayList)
			playLists.POST("/:id/tracks", h.addTrack)
			playLists.DELETE("/:id/tracks/:trackID", h.deleteTrack)
		}
	}
	router.POST("/refresh-token", h.refreshToken)
	return router
}
