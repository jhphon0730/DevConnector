package routes

import (
	"github.com/gin-gonic/gin"
)

func InitialRoutes(router *gin.Engine) {
	userRoutes(router)
}
