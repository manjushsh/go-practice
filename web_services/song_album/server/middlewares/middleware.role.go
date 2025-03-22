package middlewares

import (
	"go-practice/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware to check role access
func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user role from context
		userRole, ok := c.Get("role")
		if !ok || !services.HasAccess(userRole.(string), requiredRole) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden. You need permissions!"})
			c.Abort()
			return
		}
		c.Next()
	}
}
