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
	protected.GET("/user/all", controllers.GetAllUser)
	protected.GET("/user/:id", controllers.GetDetailUser)            // get all user
	protected.GET("/user/byrole/:id", controllers.GetUserByRole)     // get user by role
	protected.PUT("/user/edituser/:id", controllers.EditUser)        // edit user (admin only)
	protected.DELETE("/user/deleteuser/:id", controllers.DeleteUser) // delete user (admin only)
	protected.POST("/logout", controllers.Logout)
	protected.POST("/create-dprd", controllers.CreateDPRD)
	protected.GET("/index-jabatan", controllers.GetAllJabatan)
	protected.GET("/index-role", controllers.GetAllRole)

	// Route Institusi
	protected.GET("/index-institusi", controllers.GetAllInstitusi)

	// Route Program
	protected.GET("/index-program", controllers.GetAllProgram)
	protected.GET("/program/:id", controllers.DetailProgram)
	protected.POST("/program/pengajuan", controllers.PengajuanProgram)
	protected.PUT("/program/edit/:id", controllers.EditProgram)
	protected.GET("/program/accept/:id", controllers.AcceptProgram)
	protected.GET("/program/reject/:id", controllers.RejectProgram)
	protected.GET("/program/status/:status", controllers.GetProgramByStatus)
	protected.GET("/index-jenis-anggaran", controllers.GetAllJenisAnggaran)
	protected.GET("/index-kategori-penggunaan", controllers.GetAllKategoriPenggunaan)

	// Route Aduan
	protected.GET("/index-aduan", controllers.GetAllAduan)
	protected.POST("/create-aduan", controllers.CreateAduan)
	protected.GET("/count-aduan", controllers.CountAduan)
	protected.GET("/count-aduan-perbulan", controllers.CountAduanPerBulan)
	protected.GET("/count-aduan-pertahun", controllers.CountAduanPerTahun)
	protected.GET("/aduan/:id", controllers.DetailAduan)
	protected.PUT("/aduan/tanggapi/:id", controllers.TanggapiAduan)
	protected.GET("/aduan/status/:status", controllers.GetAduanByStatus)

	//Route Log
	protected.GET("/index-log", controllers.GetAllLog)

	// Route File Upload
	router.Static("/uploads", "./uploads")

	router.Run(":8000")
}
