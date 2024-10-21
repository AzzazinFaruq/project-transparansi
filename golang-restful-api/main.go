package main

import (
	"Azzazin/backend/controllers"
	middleware "Azzazin/backend/middlewares"
	"Azzazin/backend/setup"

	"github.com/gin-gonic/gin"
)

func main() {
	//Declare New Gin Route System
	router := gin.New()
	router.Use(middleware.CORSMiddleware())
	//Run Database Setup
	setup.ConnectDatabase()

	//For route that doesnt need a middleware like login, register, etc.
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	//for route that need auth with middleware
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/user", controllers.GetCurrentUser)
	protected.POST("/logout", controllers.Logout)
	protected.GET("/index-institusi", controllers.GetAllInstitusi)
	protected.GET("/index-program", controllers.GetAllProgram)
	protected.GET("/index-aduan", controllers.GetAllAduan)
	router.Run(":8000")
}
