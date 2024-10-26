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

	// Route User
	protected.GET("/user", controllers.GetCurrentUser)
	protected.GET("/user/all", controllers.GetAllUser)               // get all user
	protected.GET("/user/byrole/:id", controllers.GetUserByRole)     // get user by role
	protected.PUT("/user/edituser/:id", controllers.EditUser)        // edit user (admin only)
	protected.DELETE("/user/deleteuser/:id", controllers.DeleteUser) // delete user (admin only)
	protected.POST("/logout", controllers.Logout)

	// Route Institusi
	protected.GET("/index-institusi", controllers.GetAllInstitusi)

	// Route Program
	protected.GET("/index-program", controllers.GetAllProgram)
	protected.GET("/program/:id", controllers.DetailProgram)
	protected.POST("/program/pengajuan", controllers.PengajuanProgram)
	protected.GET("/program/accept/:id", controllers.AcceptProgram)
	protected.GET("/program/reject/:id", controllers.RejectProgram)
	protected.GET("/program/menunggu", controllers.GetProgramByStatusMenunggu)
	protected.GET("/program/disetujui", controllers.GetProgramByStatusDisetujui)
	protected.GET("/program/ditolak", controllers.GetProgramByStatusDitolak)

	// Route Aduan
	protected.GET("/index-aduan", controllers.GetAllAduan)
	protected.POST("/create-aduan", controllers.CreateAduan)
	protected.GET("/count-aduan", controllers.CountAduan)
	protected.GET("/aduan/perbulan", controllers.GetAduanPerBulan)
	protected.GET("/aduan/pertahun", controllers.GetAduanPerTahun)
	protected.GET("/aduan/:id", controllers.DetailAduan)
	protected.GET("/aduan/accept/:id", controllers.AcceptAduan)
	protected.GET("/aduan/reject/:id", controllers.RejectAduan)

	//Route Log
	protected.GET("/index-log", controllers.GetAllLog)

	router.Run(":8000")
}
