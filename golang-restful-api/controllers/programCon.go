package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllProgram(c *gin.Context) {
	var program []models.Program

	if err := setup.DB.
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("User.Jabatan").
		Preload("User.Role").
		Preload("Aspirator").
		Preload("DinasVerifikator").
		Preload("Desa").
		Preload("Kecamatan").
		Preload("Kabupaten").
		Order("created_at DESC").
		Find(&program).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	formattedPrograms := make([]gin.H, len(program))

	for i, program := range program {
		formattedPrograms[i] = gin.H{
			"id":                       program.Id,
			"nama_program":             program.NamaProgram,
			"deskripsi":                program.Deskripsi,
			"nama_institusi":           program.NamaInstitusi,
			"jenis_anggaran":           program.JenisAnggaran,
			"kategori_penggunaan":      program.KategoriPenggunaan,
			"jenis_anggaran_lain":      program.JenisAnggaranLain,
			"kategori_penggunaan_lain": program.KategoriPenggunaanLain,
			"jumlah_anggaran":          program.JumlahAnggaran,
			"aspirator":                program.Aspirator,
			"dinas_verifikator":        program.DinasVerifikator,
			"dusun":                    program.Dusun,
			"desa":                     program.Desa,
			"kecamatan":                program.Kecamatan,
			"kabupaten":                program.Kabupaten,
			"user":                     program.User,
			"status":                   program.Status,
			"foto_before":              program.FotoBefore,
			"foto_progress":            program.FotoProgress,
			"foto_after":               program.FotoAfter,
			"latitude":                 program.Latitude,
			"longitude":                program.Longitude,
			"created_at":               program.CreatedAt.Format("02-01-2006"),
			"updated_at":               program.UpdatedAt.Format("02-01-2006"),
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": formattedPrograms})
}

func GetProgramByUserId(c *gin.Context) {
	userId := c.GetInt64("user_id")

	var program []models.Program

	if err := setup.DB.
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("User.Jabatan").
		Preload("User.Role").
		Preload("Aspirator").
		Preload("DinasVerifikator").
		Preload("Desa").
		Preload("Kecamatan").
		Preload("Kabupaten").
		Where("user_id = ?", userId).
		Order("created_at DESC").
		Find(&program).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	formattedPrograms := make([]gin.H, len(program))

	for i, program := range program {
		formattedPrograms[i] = gin.H{
			"id":                       program.Id,
			"nama_program":             program.NamaProgram,
			"deskripsi":                program.Deskripsi,
			"nama_institusi":           program.NamaInstitusi,
			"jenis_anggaran":           program.JenisAnggaran,
			"kategori_penggunaan":      program.KategoriPenggunaan,
			"jenis_anggaran_lain":      program.JenisAnggaranLain,
			"kategori_penggunaan_lain": program.KategoriPenggunaanLain,
			"aspirator":                program.Aspirator,
			"dinas_verifikator":        program.DinasVerifikator,
			"dusun":                    program.Dusun,
			"desa":                     program.Desa,
			"kecamatan":                program.Kecamatan,
			"kabupaten":                program.Kabupaten,
			"user":                     program.User,
			"status":                   program.Status,
			"foto_before":              program.FotoBefore,
			"foto_progress":            program.FotoProgress,
			"foto_after":               program.FotoAfter,
			"latitude":                 program.Latitude,
			"longitude":                program.Longitude,
			"created_at":               program.CreatedAt.Format("02-01-2006"),
			"updated_at":               program.UpdatedAt.Format("02-01-2006"),
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": formattedPrograms})
}

func TambahProgram(c *gin.Context) {
	var program models.Program

	// Cek field yang wajib diisi
	namaProgram := c.PostForm("nama_program")
	if namaProgram == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama program wajib diisi"})
		return
	}

	deskripsi := c.PostForm("deskripsi")
	if deskripsi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Deskripsi wajib diisi"})
		return
	}

	// Ambil data lainnya (opsional)
	namaInstitusi := c.PostForm("nama_institusi")
	jenisAnggaranId := c.PostForm("jenis_anggaran_id")
	jumlahAnggaran := c.PostForm("jumlah_anggaran")
	kategoriPenggunaanId := c.PostForm("kategori_penggunaan_id")
	jenisAnggaranLain := c.PostForm("jenis_anggaran_lain")
	kategoriPenggunaanLain := c.PostForm("kategori_penggunaan_lain")
	aspiratorId := c.PostForm("aspirator_id")
	dinasVerifikatorId := c.PostForm("dinas_verifikator_id")
	dusun := c.PostForm("dusun")
	desaId := c.PostForm("desa_id")
	kecamatanId := c.PostForm("kecamatan_id")
	kabupatenId := c.PostForm("kabupaten_id")
	userId := c.PostForm("user_id")
	latitude := c.PostForm("latitude")
	longitude := c.PostForm("longitude")

	// Set status default ke "Draft" jika data tidak lengkap
	status := "Draft"

	// Cek kelengkapan data untuk status "Publish"
	isComplete := true

	// Parse dan validasi data opsional
	var jenisAnggaranIdInt, kategoriPenggunaanIdInt, aspiratorIdInt, dinasVerifikatorIdInt int64
	var desaIdInt, kecamatanIdInt, kabupatenIdInt, userIdInt int64
	var latitudeFloat, longitudeFloat float64
	var err error

	if jenisAnggaranId != "" {
		jenisAnggaranIdInt, err = strconv.ParseInt(jenisAnggaranId, 10, 64)
		if err != nil {
			isComplete = false
		}
	} else {
		isComplete = false
	}

	if kategoriPenggunaanId != "" {
		kategoriPenggunaanIdInt, err = strconv.ParseInt(kategoriPenggunaanId, 10, 64)
		if err != nil {
			isComplete = false
		}
	} else {
		isComplete = false
	}

	if aspiratorId != "" {
		aspiratorIdInt, err = strconv.ParseInt(aspiratorId, 10, 64)
		if err != nil {
			isComplete = false
		}
	} else {
		isComplete = false
	}

	if dinasVerifikatorId != "" {
		dinasVerifikatorIdInt, err = strconv.ParseInt(dinasVerifikatorId, 10, 64)
		if err != nil {
			isComplete = false
		}
	} else {
		isComplete = false
	}

	if desaId != "" {
		desaIdInt, err = strconv.ParseInt(desaId, 10, 64)
		if err != nil {
			isComplete = false
		}
	} else {
		isComplete = false
	}

	if kecamatanId != "" {
		kecamatanIdInt, err = strconv.ParseInt(kecamatanId, 10, 64)
		if err != nil {
			isComplete = false
		}
	} else {
		isComplete = false
	}

	if kabupatenId != "" {
		kabupatenIdInt, err = strconv.ParseInt(kabupatenId, 10, 64)
		if err != nil {
			isComplete = false
		}
	} else {
		isComplete = false
	}

	if userId != "" {
		userIdInt, err = strconv.ParseInt(userId, 10, 64)
		if err != nil {
			isComplete = false
		}
	} else {
		isComplete = false
	}

	if latitude != "" {
		latitudeFloat, err = strconv.ParseFloat(latitude, 64)
		if err != nil {
			isComplete = false
		}
	} else {
		isComplete = false
	}

	if longitude != "" {
		longitudeFloat, err = strconv.ParseFloat(longitude, 64)
		if err != nil {
			isComplete = false
		}
	} else {
		isComplete = false
	}

	// Handle foto uploads (opsional)
	fotoBefore, err := c.FormFile("foto_before")
	if err == nil {
		if !strings.HasSuffix(strings.ToLower(fotoBefore.Filename), ".jpg") &&
			!strings.HasSuffix(strings.ToLower(fotoBefore.Filename), ".jpeg") &&
			!strings.HasSuffix(strings.ToLower(fotoBefore.Filename), ".png") {
			isComplete = false
		}

		if fotoBefore.Size > 6*1024*1024 {
			isComplete = false
		}

		uploadPath := "uploads/foto-before/" + fotoBefore.Filename
		if err := c.SaveUploadedFile(fotoBefore, uploadPath); err != nil {
			isComplete = false
		}
		program.FotoBefore = uploadPath
	} else {
		isComplete = false
	}

	// Set status berdasarkan kelengkapan data
	if isComplete {
		status = "Publish"
	}

	// Buat objek program baru
	newProgram := models.Program{
		NamaProgram:            namaProgram,
		Deskripsi:              deskripsi,
		NamaInstitusi:          namaInstitusi,
		JenisAnggaranId:        jenisAnggaranIdInt,
		JumlahAnggaran:         jumlahAnggaran,
		KategoriPenggunaanId:   kategoriPenggunaanIdInt,
		JenisAnggaranLain:      jenisAnggaranLain,
		KategoriPenggunaanLain: kategoriPenggunaanLain,
		AspiratorId:            aspiratorIdInt,
		DinasVerifikatorId:     dinasVerifikatorIdInt,
		Dusun:                  dusun,
		DesaId:                 desaIdInt,
		KecamatanId:            kecamatanIdInt,
		KabupatenId:            kabupatenIdInt,
		FotoBefore:             program.FotoBefore,
		FotoProgress:           program.FotoProgress,
		FotoAfter:              program.FotoAfter,
		UserId:                 userIdInt,
		Status:                 status,
		Latitude:               latitudeFloat,
		Longitude:              longitudeFloat,
	}

	tx := setup.DB.Begin()

	if err := tx.Create(&newProgram).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan program: " + err.Error()})
		return
	}

	newLog := models.Log{
		UserId:    userIdInt,
		Aktivitas: "Menambahkan Program",
		Status:    status,
	}

	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat log aktivitas"})
		return
	}

	tx.Commit()

	setup.DB.
		Preload("JenisAnggaran").
		Preload("KategoriPenggunaan").
		Preload("Aspirator").
		Preload("DinasVerifikator").
		Preload("User").
		Preload("Desa").
		Preload("Kecamatan").
		Preload("Kabupaten").
		First(&newProgram, newProgram.Id)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Program berhasil ditambahkan dengan status " + status,
		"data":    newProgram,
	})
}

func EditProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if err := setup.DB.First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program tidak ditemukan"})
		return
	}

	namaProgram := c.PostForm("nama_program")
	deskripsi := c.PostForm("deskripsi")
	namaInstitusi := c.PostForm("nama_institusi")
	jenisAnggaranId := c.PostForm("jenis_anggaran_id")
	jumlahAnggaran := c.PostForm("jumlah_anggaran")
	kategoriPenggunaanId := c.PostForm("kategori_penggunaan_id")
	jenisAnggaranLain := c.PostForm("jenis_anggaran_lain")
	kategoriPenggunaanLain := c.PostForm("kategori_penggunaan_lain")
	aspiratorId := c.PostForm("aspirator_id")
	dinasVerifikatorId := c.PostForm("dinas_verifikator_id")
	dusun := c.PostForm("dusun")
	desaId := c.PostForm("desa_id")
	kecamatanId := c.PostForm("kecamatan_id")
	kabupatenId := c.PostForm("kabupaten_id")
	latitude := c.PostForm("latitude")
	longitude := c.PostForm("longitude")

	fotoBefore, err := c.FormFile("foto_before")
	if err == nil {
		if !strings.HasSuffix(strings.ToLower(fotoBefore.Filename), ".jpg") &&
			!strings.HasSuffix(strings.ToLower(fotoBefore.Filename), ".jpeg") &&
			!strings.HasSuffix(strings.ToLower(fotoBefore.Filename), ".png") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung"})
			return
		}

		if fotoBefore.Size > 6*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ukuran file terlalu besar"})
			return
		}

		uploadPath := "uploads/foto-before/" + fotoBefore.Filename

		if err := c.SaveUploadedFile(fotoBefore, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan foto"})
			return
		}

		program.FotoBefore = uploadPath
	}

	fotoProgress, err := c.FormFile("foto_progress")
	if err == nil {
		if !strings.HasSuffix(strings.ToLower(fotoProgress.Filename), ".jpg") &&
			!strings.HasSuffix(strings.ToLower(fotoProgress.Filename), ".jpeg") &&
			!strings.HasSuffix(strings.ToLower(fotoProgress.Filename), ".png") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung"})
			return
		}

		if fotoProgress.Size > 6*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ukuran file terlalu besar"})
			return
		}
		uploadPath := "uploads/foto-progress/" + fotoProgress.Filename

		if err := c.SaveUploadedFile(fotoProgress, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan foto"})
			return
		}

		program.FotoProgress = uploadPath
	}

	fotoAfter, err := c.FormFile("foto_after")
	if err == nil {
		if !strings.HasSuffix(strings.ToLower(fotoAfter.Filename), ".jpg") &&
			!strings.HasSuffix(strings.ToLower(fotoAfter.Filename), ".jpeg") &&
			!strings.HasSuffix(strings.ToLower(fotoAfter.Filename), ".png") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung"})
			return
		}

		if fotoAfter.Size > 6*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ukuran file terlalu besar"})
			return
		}
		uploadPath := "uploads/foto-after/" + fotoAfter.Filename

		if err := c.SaveUploadedFile(fotoAfter, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan foto"})
			return
		}

		program.FotoAfter = uploadPath
	}

	if fotoBefore == nil {
		program.FotoBefore = ""
	}
	if fotoProgress == nil {
		program.FotoProgress = ""
	}
	if fotoAfter == nil {
		program.FotoAfter = ""
	}

	tx := setup.DB.Begin()

	updateData := map[string]interface{}{
		"nama_program":             namaProgram,
		"deskripsi":                deskripsi,
		"nama_institusi":           namaInstitusi,
		"jenis_anggaran_id":        jenisAnggaranId,
		"jumlah_anggaran":          jumlahAnggaran,
		"kategori_penggunaan_id":   kategoriPenggunaanId,
		"jenis_anggaran_lain":      jenisAnggaranLain,
		"kategori_penggunaan_lain": kategoriPenggunaanLain,
		"aspirator_id":             aspiratorId,
		"dinas_verifikator_id":     dinasVerifikatorId,
		"dusun":                    dusun,
		"desa_id":                  desaId,
		"kecamatan_id":             kecamatanId,
		"kabupaten_id":             kabupatenId,
		"foto_before":              program.FotoBefore,
		"foto_progress":            program.FotoProgress,
		"foto_after":               program.FotoAfter,
		"latitude":                 latitude,
		"longitude":                longitude,
	}

	// Hanya update field yang tidak kosong
	for key, value := range updateData {
		if value != "" {
				tx.Model(&program).UpdateColumn(key, value)
		} else if value == "" {
			tx.Model(&program).UpdateColumn(key, "")
		}
	}

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

	setup.DB.
		Preload("JenisAnggaran").
		Preload("KategoriPenggunaan").
		Preload("Aspirator").
		Preload("DinasVerifikator").
		Preload("User").
		First(&program, program.Id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Program berhasil diupdate",
		"data":    program,
	})
}

func GetProgramByStatus(c *gin.Context) {
	status := c.Param("status")

	validStatus := map[string]bool{
		"Publish": true,
		"Draft":   true,
	}

	if !validStatus[status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status tidak valid"})
		return
	}

	var programs []models.Program

	if err := setup.DB.
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("Aspirator").
		Preload("DinasVerifikator").
		Preload("User.Jabatan").
		Preload("User.Role").
		Where("status = ?", status).
		Order("created_at DESC").
		Find(&programs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	formattedPrograms := make([]gin.H, len(programs))

	for i, program := range programs {
		formattedPrograms[i] = gin.H{
			"id":                       program.Id,
			"nama_program":             program.NamaProgram,
			"deskripsi":                program.Deskripsi,
			"nama_institusi":           program.NamaInstitusi,
			"jenis_anggaran":           program.JenisAnggaran,
			"jumlah_anggaran":          program.JumlahAnggaran,
			"kategori_penggunaan":      program.KategoriPenggunaan,
			"jenis_anggaran_lain":      program.JenisAnggaranLain,
			"kategori_penggunaan_lain": program.KategoriPenggunaanLain,
			"aspirator":                program.Aspirator,
			"dinas_verifikator":        program.DinasVerifikator,
			"dusun":                    program.Dusun,
			"user":                     program.User,
			"status":                   program.Status,
			"foto_before":              program.FotoBefore,
			"foto_progress":            program.FotoProgress,
			"foto_after":               program.FotoAfter,
			"latitude":                 program.Latitude,
			"longitude":                program.Longitude,
			"created_at":               program.CreatedAt.Format("02-01-2006"),
			"updated_at":               program.UpdatedAt.Format("02-01-2006"),
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

	result := setup.DB.
		Preload("Desa").
		Preload("Kecamatan").
		Preload("Kabupaten").
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("Aspirator").
		Preload("DinasVerifikator").
		Preload("User.Jabatan").
		Preload("User.Role").
		First(&program, id)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Program dengan ID " + id + " tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan: " + result.Error.Error()})
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
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("Aspirator").
		Preload("DinasVerifikator").
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
			"id":                       program.Id,
			"nama_program":             program.NamaProgram,
			"deskripsi":                program.Deskripsi,
			"nama_institusi":           program.NamaInstitusi,
			"jenis_anggaran":           program.JenisAnggaran,
			"jumlah_anggaran":          program.JumlahAnggaran,
			"kategori_penggunaan":      program.KategoriPenggunaan,
			"jenis_anggaran_lain":      program.JenisAnggaranLain,
			"kategori_penggunaan_lain": program.KategoriPenggunaanLain,
			"aspirator":                program.Aspirator,
			"dinas_verifikator":        program.DinasVerifikator,
			"dusun":                    program.Dusun,
			"user":                     program.User,
			"status":                   program.Status,
			"foto_before":              program.FotoBefore,
			"foto_progress":            program.FotoProgress,
			"foto_after":               program.FotoAfter,
			"latitude":                 program.Latitude,
			"longitude":                program.Longitude,
			"created_at":               program.CreatedAt.Format("02-01-2006"),
			"updated_at":               program.UpdatedAt.Format("02-01-2006"),
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
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("Aspirator").
		Preload("DinasVerifikator").
		Preload("User.Jabatan").
		Preload("User.Role").
		Where("status = ?", "Publish").
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

func GetProgramByDaerah(c *gin.Context) {
	queryDesa := c.Query("desa")
	queryKecamatan := c.Query("kecamatan")
	queryKabupaten := c.Query("kabupaten")

	if queryDesa == "" && queryKecamatan == "" && queryKabupaten == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pencarian tidak boleh kosong"})
		return
	}

	var programs []models.Program

	if err := setup.DB.
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("Aspirator").
		Preload("DinasVerifikator").
		Preload("User.Jabatan").
		Preload("User.Role").
		Where("desa_id = ?", queryDesa).
		Where("kecamatan_id = ?", queryKecamatan).
		Where("kabupaten_id = ?", queryKabupaten).
		Find(&programs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(programs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Tidak ada program yang ditemukan",
			"data":    []interface{}{},
		})
		return
	}

	formattedPrograms := make([]gin.H, len(programs))
	for i, program := range programs {
		formattedPrograms[i] = gin.H{
			"id":                       program.Id,
			"nama_program":             program.NamaProgram,
			"deskripsi":                program.Deskripsi,
			"nama_institusi":           program.NamaInstitusi,
			"jenis_anggaran":           program.JenisAnggaran,
			"jumlah_anggaran":          program.JumlahAnggaran,
			"kategori_penggunaan":      program.KategoriPenggunaan,
			"jenis_anggaran_lain":      program.JenisAnggaranLain,
			"kategori_penggunaan_lain": program.KategoriPenggunaanLain,
			"aspirator":                program.Aspirator,
			"dinas_verifikator":        program.DinasVerifikator,
			"dusun":                    program.Dusun,
			"user":                     program.User,
			"status":                   program.Status,
			"foto_before":              program.FotoBefore,
			"foto_progress":            program.FotoProgress,
			"foto_after":               program.FotoAfter,
			"latitude":                 program.Latitude,
			"longitude":                program.Longitude,
			"created_at":               program.CreatedAt.Format("02-01-2006"),
			"updated_at":               program.UpdatedAt.Format("02-01-2006"),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Program ditemukan",
		"total":   len(programs),
		"data":    formattedPrograms,
	})
}

func DraftProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if err := setup.DB.First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program tidak ditemukan"})
		return
	}

	tx := setup.DB.Begin()

	program.Status = "Draft"
	if err := tx.Save(&program).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengubah status program"})
		return
	}

	newLog := models.Log{
		UserId:    program.UserId,
		Aktivitas: "Draft Program",
		Status:    "Draft",
	}

	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat log aktivitas"})
		return
	}

	tx.Commit()

	setup.DB.Preload("Institusi").
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("Aspirator").
		Preload("DinasVerifikator").
		Preload("User").
		First(&program, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Program berhasil didraft",
		"data":    program,
	})
}

func PublishProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if err := setup.DB.First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program tidak ditemukan"})
		return
	}

	tx := setup.DB.Begin()

	program.Status = "Publish"
	if err := tx.Save(&program).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengubah status program"})
		return
	}

	newLog := models.Log{
		UserId:    program.UserId,
		Aktivitas: "Publish Program",
		Status:    "Publish",
	}

	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat log aktivitas"})
		return
	}

	tx.Commit()

	setup.DB.Preload("Institusi").
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("Aspirator").
		Preload("DinasVerifikator").
		Preload("User").
		First(&program, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Program berhasil dipublish",
		"data":    program,
	})
}
