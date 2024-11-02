package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllProgram(c *gin.Context) {
	var program []models.Program

	if err := setup.DB.
		Preload("Institusi").
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("User.Jabatan").
		Preload("User.Role").
		Find(&program).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	formattedPrograms := make([]gin.H, len(program))

	for i, program := range program {
		formattedPrograms[i] = gin.H{
			"id":                  program.Id,
			"nama_program":        program.NamaProgram,
			"deskripsi":           program.Deskripsi,
			"institusi":           program.Institusi,
			"jenis_anggaran":      program.JenisAnggaran,
			"kategori_penggunaan": program.KategoriPenggunaan,
			"dusun":               program.Dusun,
			"user":                program.User,
			"status":              program.Status,
			"foto_before":         program.FotoBefore,
			"foto_progress":       program.FotoProgress,
			"foto_after":          program.FotoAfter,
			"created_at":          program.CreatedAt.Format("02-01-2006"),
			"updated_at":          program.UpdatedAt.Format("02-01-2006"),
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": formattedPrograms})
}

func PengajuanProgram(c *gin.Context) {
	var input models.Program

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid"})
		return
	}

	newProgram := models.Program{
		NamaProgram:          input.NamaProgram,
		Deskripsi:            input.Deskripsi,
		InstitusiId:          input.InstitusiId,
		JenisAnggaranId:      input.JenisAnggaranId,
		JumlahAnggaran:       input.JumlahAnggaran,
		KategoriPenggunaanId: input.KategoriPenggunaanId,
		Dusun:                input.Dusun,
		DesaId:               input.DesaId,
		KecamatanId:          input.KecamatanId,
		KabupatenId:          input.KabupatenId,
		UserId:               input.UserId,
		Status:               "Menunggu",
	}

	tx := setup.DB.Begin()

	if err := tx.Create(&newProgram).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengajukan program baru"})
		return
	}

	newLog := models.Log{
		UserId:    input.UserId,
		Aktivitas: "Pengajuan Program",
		Status:    "Menunggu",
	}

	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat log aktivitas"})
		return
	}

	tx.Commit()

	setup.DB.Preload("Institusi").
		Preload("JenisAnggaran").
		Preload("KategoriPenggunaan").
		Preload("User").
		First(&newProgram, newProgram.Id)

	c.JSON(http.StatusCreated, gin.H{"message": "Program berhasil diajukan", "data": newProgram})
}

func EditProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if err := setup.DB.First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program tidak ditemukan"})
		return
	}

	// Bind form data
	namaProgram := c.PostForm("nama_program")
	deskripsi := c.PostForm("deskripsi")
	institusiId := c.PostForm("institusi_id")
	jenisAnggaranId := c.PostForm("jenis_anggaran_id")
	jumlahAnggaran := c.PostForm("jumlah_anggaran")
	kategoriPenggunaanId := c.PostForm("kategori_penggunaan_id")
	dusun := c.PostForm("dusun")
	desaId := c.PostForm("desa_id")
	kecamatanId := c.PostForm("kecamatan_id")
	kabupatenId := c.PostForm("kabupaten_id")

	// Handle foto before
	fotoBefore, err := c.FormFile("foto_before")
	if err == nil {
		// Validasi tipe file
		if !strings.HasSuffix(strings.ToLower(fotoBefore.Filename), ".jpg") &&
			!strings.HasSuffix(strings.ToLower(fotoBefore.Filename), ".jpeg") &&
			!strings.HasSuffix(strings.ToLower(fotoBefore.Filename), ".png") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung"})
			return
		}

		// Validasi ukuran file (max 6MB)
		if fotoBefore.Size > 6*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ukuran file terlalu besar"})
			return
		}

		// Simpan file di folder uploads/foto-before
		uploadPath := "uploads/foto-before/" + fotoBefore.Filename

		// Simpan file
		if err := c.SaveUploadedFile(fotoBefore, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan foto"})
			return
		}

		// Simpan path ke database
		program.FotoBefore = uploadPath
	}

	// Handle foto progress
	fotoProgress, err := c.FormFile("foto_progress")
	if err == nil {
		// Validasi tipe file
		if !strings.HasSuffix(strings.ToLower(fotoProgress.Filename), ".jpg") &&
			!strings.HasSuffix(strings.ToLower(fotoProgress.Filename), ".jpeg") &&
			!strings.HasSuffix(strings.ToLower(fotoProgress.Filename), ".png") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung"})
			return
		}

		// Validasi ukuran file (max 6MB)
		if fotoProgress.Size > 6*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ukuran file terlalu besar"})
			return
		}
		// Simpan file di folder uploads/foto-progress
		uploadPath := "uploads/foto-progress/" + fotoProgress.Filename

		// Simpan file
		if err := c.SaveUploadedFile(fotoProgress, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan foto"})
			return
		}

		// Simpan path ke database
		program.FotoProgress = uploadPath
	}

	// Handle foto after
	fotoAfter, err := c.FormFile("foto_after")
	if err == nil {
		// Validasi tipe file
		if !strings.HasSuffix(strings.ToLower(fotoAfter.Filename), ".jpg") &&
			!strings.HasSuffix(strings.ToLower(fotoAfter.Filename), ".jpeg") &&
			!strings.HasSuffix(strings.ToLower(fotoAfter.Filename), ".png") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung"})
			return
		}

		// Validasi ukuran file (max 6MB)
		if fotoAfter.Size > 6*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ukuran file terlalu besar"})
			return
		}
		// Simpan file di folder uploads/foto-after
		uploadPath := "uploads/foto-after/" + fotoAfter.Filename

		// Simpan file
		if err := c.SaveUploadedFile(fotoAfter, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan foto"})
			return
		}

		// Simpan path ke database
		program.FotoAfter = uploadPath
	}

	// Update data program
	tx := setup.DB.Begin()

	updateData := map[string]interface{}{
		"nama_program":           namaProgram,
		"deskripsi":              deskripsi,
		"institusi_id":           institusiId,
		"jenis_anggaran_id":      jenisAnggaranId,
		"jumlah_anggaran":        jumlahAnggaran,
		"kategori_penggunaan_id": kategoriPenggunaanId,
		"dusun":                  dusun,
		"desa_id":                desaId,
		"kecamatan_id":           kecamatanId,
		"kabupaten_id":           kabupatenId,
		"foto_before":            program.FotoBefore,
		"foto_progress":          program.FotoProgress,
		"foto_after":             program.FotoAfter,
	}

	// Hanya update field yang tidak kosong
	for key, value := range updateData {
		if value != "" {
			tx.Model(&program).UpdateColumn(key, value)
		}
	}

	// Catat log aktivitas
	newLog := models.Log{
		UserId:    program.UserId,
		Aktivitas: "Edit Program",
		Status:    program.Status,
	}

	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat log aktivitas"})
		return
	}

	tx.Commit()

	// Ambil data terbaru
	setup.DB.Preload("Institusi").
		Preload("JenisAnggaran").
		Preload("KategoriPenggunaan").
		Preload("User").
		First(&program, program.Id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Program berhasil diupdate",
		"data":    program,
	})
}

func AcceptProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if err := setup.DB.First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program tidak ditemukan"})
		return
	}

	program.Status = "Dalam Proses"

	tx := setup.DB.Begin()

	if err := tx.Save(&program).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyetujui program"})
		return
	}

	newLog := models.Log{
		UserId:    program.UserId,
		Aktivitas: "Penyetujuan Program",
		Status:    "Dalam Proses",
	}

	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat log aktivitas"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Program berhasil disetujui", "data": program})
}

func RejectProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if err := setup.DB.Find(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program tidak ditemukan"})
		return
	}

	program.Status = "Ditolak"

	tx := setup.DB.Begin()

	if err := tx.Save(&program).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menolak program"})
		return
	}

	newLog := models.Log{
		UserId:    program.UserId,
		Aktivitas: "Penolakan Program",
		Status:    "Ditolak",
	}

	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat log aktivitas"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Program berhasil ditolak", "data": program})
}

func GetProgramByStatus(c *gin.Context) {
	status := c.Param("status")

	validStatus := map[string]bool{
		"Menunggu":     true,
		"Dalam Proses": true,
		"Selesai":      true,
		"Ditolak":      true,
	}

	if !validStatus[status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status tidak valid"})
		return
	}

	var programs []models.Program

	if err := setup.DB.
		Preload("Institusi").
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("User.Jabatan").
		Preload("User.Role").
		Where("status = ?", status).
		Find(&programs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	formattedPrograms := make([]gin.H, len(programs))

	for i, program := range programs {
		formattedPrograms[i] = gin.H{
			"id":                  program.Id,
			"nama_program":        program.NamaProgram,
			"institusi":           program.Institusi,
			"jenis_anggaran":      program.JenisAnggaran,
			"jumlah_anggaran":     program.JumlahAnggaran,
			"kategori_penggunaan": program.KategoriPenggunaan,
			"dusun":               program.Dusun,
			"user":                program.User,
			"status":              program.Status,
			"foto_before":         program.FotoBefore,
			"foto_progress":       program.FotoProgress,
			"foto_after":          program.FotoAfter,
			"created_at":          program.CreatedAt.Format("02-01-2006"),
			"updated_at":          program.UpdatedAt.Format("02-01-2006"),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data":   formattedPrograms,
	})
}

func DetailProgram(c *gin.Context) {
	id := c.Param("id")

	var program models.Program

	if err := setup.DB.
		Preload("Desa").
		Preload("Kecamatan").
		Preload("Kabupaten").
		Preload("Institusi").
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("User.Jabatan").
		Preload("User.Role").
		First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": program})
}

func SearchProgram(c *gin.Context) {
	// Ambil parameter pencarian dari query string
	query := c.Query("nama")

	// Jika query kosong
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pencarian tidak boleh kosong"})
		return
	}

	var programs []models.Program

	// Gunakan LIKE untuk pencarian partial match
	if err := setup.DB.
		Preload("Institusi").
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("User.Jabatan").
		Preload("User.Role").
		Where("nama_program LIKE ?", "%"+query+"%").
		Find(&programs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Jika tidak ada hasil
	if len(programs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Tidak ada program yang ditemukan",
			"data":    []interface{}{},
		})
		return
	}

	// Format hasil pencarian
	formattedPrograms := make([]gin.H, len(programs))
	for i, program := range programs {
		formattedPrograms[i] = gin.H{
			"id":                  program.Id,
			"nama_program":        program.NamaProgram,
			"deskripsi":           program.Deskripsi,
			"institusi":           program.Institusi,
			"jenis_anggaran":      program.JenisAnggaran,
			"jumlah_anggaran":     program.JumlahAnggaran,
			"kategori_penggunaan": program.KategoriPenggunaan,
			"dusun":               program.Dusun,
			"user":                program.User,
			"status":              program.Status,
			"foto_before":         program.FotoBefore,
			"foto_progress":       program.FotoProgress,
			"foto_after":          program.FotoAfter,
			"created_at":          program.CreatedAt.Format("02-01-2006"),
			"updated_at":          program.UpdatedAt.Format("02-01-2006"),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Program ditemukan",
		"total":   len(programs),
		"data":    formattedPrograms,
	})
}

func GetProgramLandingPage(c *gin.Context) {
	var program []models.Program

	if err := setup.DB.
		Preload("Institusi").
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("User.Jabatan").
		Preload("User.Role").
		Find(&program).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	formattedPrograms := make([]gin.H, len(program))

	for i, program := range program {
		formattedPrograms[i] = gin.H{
			"id":           program.Id,
			"nama_program": program.NamaProgram,
			"deskripsi":    program.Deskripsi,
			"foto_before":  program.FotoBefore,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": formattedPrograms})
}

func SelesaikanProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if err := setup.DB.First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program tidak ditemukan"})
		return
	}

	program.Status = "Selesai"

	tx := setup.DB.Begin()

	if err := tx.Save(&program).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyelesaikan program"})
		return
	}

	newLog := models.Log{
		UserId:    program.UserId,
		Aktivitas: "Penyelesaian Program",
		Status:    "Selesai",
	}

	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat log aktivitas"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Program berhasil diselesaikan", "data": program})
}

func GetProgramByDaerah(c *gin.Context) {
	// Ambil parameter pencarian dari query string
	queryDesa := c.Query("desa")
	queryKecamatan := c.Query("kecamatan")
	queryKabupaten := c.Query("kabupaten")

	// Jika query kosong
	if queryDesa == "" && queryKecamatan == "" && queryKabupaten == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pencarian tidak boleh kosong"})
		return
	}

	var programs []models.Program

	// Gunakan LIKE untuk pencarian partial match
	if err := setup.DB.
		Preload("Institusi").
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("User.Jabatan").
		Preload("User.Role").
		Where("desa_id = ?", queryDesa).
		Where("kecamatan_id = ?", queryKecamatan).
		Where("kabupaten_id = ?", queryKabupaten).
		Find(&programs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Jika tidak ada hasil
	if len(programs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Tidak ada program yang ditemukan",
			"data":    []interface{}{},
		})
		return
	}

	// Format hasil pencarian
	formattedPrograms := make([]gin.H, len(programs))
	for i, program := range programs {
		formattedPrograms[i] = gin.H{
			"id":                  program.Id,
			"nama_program":        program.NamaProgram,
			"deskripsi":           program.Deskripsi,
			"institusi":           program.Institusi,
			"jenis_anggaran":      program.JenisAnggaran,
			"jumlah_anggaran":     program.JumlahAnggaran,
			"kategori_penggunaan": program.KategoriPenggunaan,
			"dusun":               program.Dusun,
			"user":                program.User,
			"status":              program.Status,
			"foto_before":         program.FotoBefore,
			"foto_progress":       program.FotoProgress,
			"foto_after":          program.FotoAfter,
			"created_at":          program.CreatedAt.Format("02-01-2006"),
			"updated_at":          program.UpdatedAt.Format("02-01-2006"),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Program ditemukan",
		"total":   len(programs),
		"data":    formattedPrograms,
	})
}
