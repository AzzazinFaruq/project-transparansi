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
	router.GET("/program", controllers.GetProgramLandingPage)

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
	protected.GET("/index-role", controllers.GetAllRole)
	protected.GET("/user/username", controllers.GetUserByUsername)
	
	// Route Jabatan
	protected.GET("/index-jabatan", controllers.GetAllJabatan)
	protected.POST("/create-jabatan", controllers.CreateJabatan)
	protected.GET("/jabatan/search", controllers.GetJabatanByNamaJabatan)
	protected.DELETE("/jabatan/:id", controllers.DeleteJabatan)

	// Route Program
	protected.GET("/index-program", controllers.GetAllProgram)
	protected.GET("/program/user", controllers.GetProgramByUserId)
	protected.GET("/program/:id", controllers.DetailProgram)
	protected.POST("/program/tambah", controllers.TambahProgram)
	protected.PUT("/program/edit/:id", controllers.EditProgram)
	protected.GET("/program/draft/:id", controllers.DraftProgram)
	protected.GET("/program/status/:status", controllers.GetProgramByStatus)
	protected.GET("/index-jenis-anggaran", controllers.GetAllJenisAnggaran)
	protected.GET("/index-kategori-penggunaan", controllers.GetAllKategoriPenggunaan)
	protected.GET("/program/search", controllers.SearchProgram)
	protected.GET("/program/daerah", controllers.GetProgramByDaerah)


  protected.POST("/kategori-penggunaan", controllers.CreateKategoriPenggunaan)
  protected.DELETE("/kategori-penggunaan/:id", controllers.DeleteKategoriPenggunaan)
  protected.POST("/jenis-anggaran", controllers.CreateJenisAnggaran)
  protected.DELETE("/jenis-anggaran/:id", controllers.DeleteJenisAnggaran)


	// Route Aduan
	protected.GET("/index-aduan", controllers.GetAllAduan)
	protected.POST("/create-aduan", controllers.CreateAduan)
	protected.GET("/count-aduan", controllers.CountAduan)
	protected.GET("/count-aduan-perbulan", controllers.CountAduanPerBulan)
	protected.GET("/count-aduan-pertahun", controllers.CountAduanPerTahun)
	protected.GET("/aduan/:id", controllers.DetailAduan)
	protected.PUT("/aduan/tanggapi/:id", controllers.TanggapiAduan)
	protected.GET("/aduan/status/:status", controllers.GetAduanByStatus)
	protected.GET("/aduan/aduan-by-program/:program_id", controllers.GetAduanByProgramId)

	// Route Daerah
	protected.GET("/index-kabupaten", controllers.GetAllKabupaten)
	protected.GET("/index-kecamatan", controllers.GetAllKecamatan)
	protected.GET("/index-desa", controllers.GetAllDesa)
	protected.GET("/kecamatan/:id", controllers.GetKecamatanByKabupatenId)
	protected.GET("/desa/:id", controllers.GetDesaByKecamatanId)

	// Route Aspirator
	protected.GET("/index-aspirator", controllers.GetAllAspirator)
	protected.GET("/aspirator/search", controllers.GetAspiratorByNamaAspirator)
	protected.POST("/create-aspirator", controllers.CreateAspirator)
	protected.DELETE("/aspirator/:id", controllers.DeleteAspirator)

	// Route Dinas Verifikator
	protected.GET("/index-dinas-verifikator", controllers.GetAllDinasVerifikator)
	protected.GET("/dinas-verifikator/search", controllers.GetDinasVerifikatorByNama)
	protected.POST("/create-dinas-verifikator", controllers.CreateDinasVerifikator)
	protected.DELETE("/dinas-verifikator/:id", controllers.DeleteDinasVerifikator)

	//Route Log
	protected.GET("/index-log", controllers.GetAllLog)

	// Route File Upload
	router.Static("/uploads", "./uploads")

	router.Run(":8000")
}
