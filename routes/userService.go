package routes 

import (
	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.Engine) {
	routes := router.Group("/api")
	{
		userRoute := routes.Group("/users")
		userRoute.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "User Works",
			})
		})
	}
}
