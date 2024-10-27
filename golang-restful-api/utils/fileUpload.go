package utils

import (
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context, file *multipart.FileHeader, folder string) (string, error) {
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return "", err
	}

	filename := filepath.Base(file.Filename)
	finalPath := filepath.Join(folder, filename)

	if err := c.SaveUploadedFile(file, finalPath); err != nil {
		return "", err
	}

	return finalPath, nil
}
