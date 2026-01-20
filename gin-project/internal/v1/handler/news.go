package handler

import (
	"gin-project/internal/utils"
	"os"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
}

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

type CreateNewsRequest struct {
	Title string `form:"title" binding:"required,min=5,max=100"`
}

func (h *NewsHandler) CreateNews(ctx *gin.Context) {
	var params CreateNewsRequest
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(400, utils.HandleValidationError(err))
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(400, gin.H{"error": gin.H{"image": "Image is required"}})
		return
	}

	dir := "./uploads/"

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create upload directory"})
		return
	}

	dst := dir + file.Filename
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to save uploaded file"})
		return
	}

	ctx.JSON(201, gin.H{"message": "News created", "title": params.Title})
}
