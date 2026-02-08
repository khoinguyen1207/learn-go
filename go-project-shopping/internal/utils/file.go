package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var allowedExts = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}
var mimeTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
}
var MAX_FILE_SIZE int64 = 3 * 1024 * 1024 // 3MB

func HandleUploadImage(fileHeader *multipart.FileHeader, uploadDir string) (string, error) {
	// Validate file extension
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowedExts[ext] {
		return "", fmt.Errorf("file extension %s is not allowed", ext)
	}

	fileName := uuid.New().String() + ext
	dst := filepath.Join(uploadDir, fileName)

	// Validate size
	if fileHeader.Size > MAX_FILE_SIZE {
		return "", fmt.Errorf("file size exceeds the maximum limit of %d bytes", MAX_FILE_SIZE)
	}

	//Validate mime type
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	mimeType := http.DetectContentType(buffer)
	if !mimeTypes[mimeType] {
		return "", fmt.Errorf("file mime type %s is not allowed", mimeType)
	}

	err = InitUploadDir(uploadDir)
	if err != nil {
		return "", err
	}

	err = SaveUploadedFile(fileHeader, dst)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func InitUploadDir(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create upload directory: %v", err)
	}
	return nil
}

func SaveUploadedFile(fileHeader *multipart.FileHeader, destination string) error {
	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
