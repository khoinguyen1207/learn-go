package handler

import (
	"gin-project/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

type GetProductParams struct {
	Name string `form:"name" binding:"required,min=3,max=50"`
	Page int    `form:"page" binding:"omitempty,gte=1"`
	Date string `form:"date" binding:"omitempty,datetime=2006-01-02"`
	// Limit int    `form:"limit" binding:"gt=1,max=100"`
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (h *ProductHandler) GetProduct(ctx *gin.Context) {
	var params GetProductParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Date == "" {
		params.Date = time.Now().Format("2006-01-02")
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "GetProduct called", "params": params})
}
