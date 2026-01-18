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
}

type ProductImage struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
	URL  string `json:"url" binding:"required,url"`
}

type CreateProductRequest struct {
	Name     string       `json:"name" binding:"required,min=3,max=50"`
	Price    int          `json:"price" binding:"required,gt=100"`
	IsActive *bool        `json:"is_active"`
	Image    ProductImage `json:"image" binding:"required"`
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

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var body CreateProductRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	if body.IsActive == nil {
		defaultActive := true
		body.IsActive = &defaultActive
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Product created", "product": body})
}
