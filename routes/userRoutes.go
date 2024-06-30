package routes 

import (
	"github.com/gin-gonic/gin"
	"jhphon0730/DevConnector/controllers"
	"jhphon0730/DevConnector/services"
)

var (
	userService services.UserService = services.NewUserService()
	userController controllers.UserController = controllers.NewUserController(userService)
)

func userRoutes(router *gin.Engine) {
	routes := router.Group("/api")
	{
		userRoute := routes.Group("/users")
		{
			userRoute.GET("/:user_id", userController.GetUser)
			userRoute.POST("/", userController.CreateUser)

			experienceRoute := userRoute.Group("/experience")
			{
				experienceRoute.POST("/:user_id", userController.CreateUserExperience)
			}
		}
	}
}
