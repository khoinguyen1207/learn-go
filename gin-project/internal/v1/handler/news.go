package handler

import (
	"gin-project/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
}

var UPLOAD_DIR = "./uploads/"
var URL = "http://localhost:8080"

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

type CreateNewsRequest struct {
	Title string `form:"title" binding:"required,min=5,max=100"`
}

func (h *NewsHandler) CreateNews(ctx *gin.Context) {
	var params CreateNewsRequest
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadGateway, utils.HandleValidationError(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "News created", "title": params.Title})
}

func (h *NewsHandler) UploadNewsImage(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	filename, err := utils.HandleUploadImage(file, UPLOAD_DIR)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Upload successful", "filename": filename, "url": "/uploads/" + filename})
}

func (h *NewsHandler) UploadMultipleNewsImages(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Images are required"})
		return
	}

	files := form.File["images"]

	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "At least one image is required"})
		return
	}

	successfulUploads := []string{}
	failedUploads := []map[string]string{}

	for _, file := range files {
		fileName, err := utils.HandleUploadImage(file, UPLOAD_DIR)
		if err != nil {
			failedUploads = append(failedUploads, map[string]string{
				"filename": file.Filename,
				"error":    err.Error(),
			})
			continue
		}
		imageUrl := URL + "/images/" + fileName
		successfulUploads = append(successfulUploads, imageUrl)

	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":            "Upload multiple news images endpoint",
		"successful_uploads": successfulUploads,
		"failed_uploads":     failedUploads,
	})
}
