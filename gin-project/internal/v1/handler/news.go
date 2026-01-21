package handler

import (
	"gin-project/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
}

var UPLOAD_DIR = "./uploads/"

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

	filename, err := utils.UploadSingleFile(file, UPLOAD_DIR)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Upload successful", "filename": filename, "url": "/uploads/" + filename})
}
