package handlers

import (
	"PlaylistsSynchronizer/pkg/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
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
		}
		userGroups := api.Group("/user-groups")
		{
			userGroups.POST("/", h.createUserGroup)
			userGroups.GET("/", h.getAllUserGroups)
			userGroups.GET("/:id", h.getUserGroupById)
			userGroups.PUT("/:id", h.updateUserGroup)
			userGroups.DELETE("/:id", h.deleteUserGroup)
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
			playLists.POST("/", h.createPlayList)
			playLists.GET("/", h.getAllPlayList)
			playLists.GET("/:id", h.getPlayListById)
			playLists.PUT("/:id", h.updatePlayList)
			playLists.DELETE("/:id", h.deletePlayList)
		}
	}
	return router
}
