package routes

import (
	"go-practice/controllers"
	"go-practice/middlewares"
	"go-practice/services"

	"github.com/gin-gonic/gin"
)

func RegisterAlbumRoutes(r *gin.Engine) {
	albumGroup := r.Group("/albums")
	albumGroup.Use(middlewares.AuthMiddleware())
	albumGroup.GET("/", middlewares.RoleMiddleware(services.RoleUser), controllers.GetAlbums)
	albumGroup.GET("/:id", controllers.GetAlbumByID)
	albumGroup.POST("/", middlewares.RoleMiddleware(services.RoleAdmin), controllers.PostAlbums)
	albumGroup.PUT("/:id", middlewares.RoleMiddleware(services.RoleAdmin), controllers.UpdateAlbum)
	albumGroup.DELETE("/:id", middlewares.RoleMiddleware(services.RoleAdmin), controllers.DeleteAlbum)
}
