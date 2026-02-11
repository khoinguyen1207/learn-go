package dto

import (
	"net/http"
	"project-shopping/internal/utils"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(ctx *gin.Context, err error) {
	if appErr, ok := err.(*utils.AppError); ok {
		httpStatus := httpStatusFromErrorCode(appErr.Code)

		response := gin.H{
			"message": appErr.Message,
			"code":    appErr.Code,
		}
		if appErr.Err != nil {
			response["error"] = appErr.Err.Error()
		}

		ctx.JSON(httpStatus, response)
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal Server Error",
		"code":    utils.CodeInternalServerError,
		"error":   err.Error(),
	})
}

func SuccessResponse(ctx *gin.Context, message string, data any, status ...int) {
	if len(status) > 0 {
		ctx.JSON(status[0], gin.H{
			"message": message,
			"data":    data,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}

func ValidationResponse(ctx *gin.Context, err any) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "Validation Error",
		"code":    utils.CodeValidationError,
		"errors":  err,
	})
}

func httpStatusFromErrorCode(code utils.ErrorResponseCode) int {
	switch code {
	case utils.CodeBadRequest, utils.CodeValidationError:
		return http.StatusBadRequest
	case utils.CodeUnauthorized:
		return http.StatusUnauthorized
	case utils.CodeForbidden:
		return http.StatusForbidden
	case utils.CodeNotFound:
		return http.StatusNotFound
	case utils.CodeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
