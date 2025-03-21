package routes

import (
	"go-practice/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	authGroup.POST("/login", controllers.Login)
	authGroup.POST("/register", controllers.Register)
}
