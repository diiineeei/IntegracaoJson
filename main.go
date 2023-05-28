package main

import (
	"github.com/diiineeei/IntegracaoJson/controllers"
	"github.com/diiineeei/IntegracaoJson/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.Connect("admin:admin@tcp(localhost:3306)/admin?parseTime=true")
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
	}
	return router
}
