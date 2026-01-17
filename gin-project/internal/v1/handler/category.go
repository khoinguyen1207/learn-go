package handler

import (
	"gin-project/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
}

type GetCategoryParams struct {
	Category string `uri:"category" binding:"oneof=books electronics clothing"`
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (h *CategoryHandler) GetCategory(ctx *gin.Context) {
	var params GetCategoryParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully retrieved category", "category": params.Category})
}
