package main 

import (
	"github.com/gin-gonic/gin"
	"jhphon0730/DevConnector/routes"
	"jhphon0730/DevConnector/database"
)

func main() {
	// Init Database
	database.InitialDatabase()

	// Init Gin Router Engine 
	router := gin.Default()

	// Init Routes  
	routes.InitialRoutes(router)

	// Run Server
	router.Run(":8080")
}
