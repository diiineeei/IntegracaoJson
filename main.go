package main

import (
	"github.com/diiineeei/IntegracaoJson/controllers"
	"github.com/diiineeei/IntegracaoJson/database"
	"github.com/diiineeei/IntegracaoJson/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	// Initialize Database
	database.ConnectMySQL("admin:admin@tcp(localhost:3306)/admin?parseTime=true")
	database.ConnectPostgres("postgres://admin:admin@localhost:5432/admin")
	database.Migrate()
}
func main() {
	// Initialize Router
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.POST("/contact/register", controllers.RegisterContacts)
			secured.POST("/contact/register/list", controllers.RegisterContactsList)
		}
	}
	router.Run(":8080")
}
