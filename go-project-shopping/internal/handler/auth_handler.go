package handler

import (
	"project-shopping/internal/dto"
	"project-shopping/internal/service"
	"project-shopping/internal/validation"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (ah *AuthHandler) Login(ctx *gin.Context) {
	var params dto.LoginRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	accessToken, refreshToken, err := ah.service.Login(ctx, params.Email, params.Password)
	if err != nil {
		dto.ErrorResponse(ctx, err)
		return
	}

	response := dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	dto.SuccessResponse(ctx, "Login successfully", response)

}

func (ah *AuthHandler) Register(ctx *gin.Context) {

}

func (ah *AuthHandler) Logout(ctx *gin.Context) {
	var params dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}
	err := ah.service.Logout(ctx.Request.Context(), params.RefreshToken)
	if err != nil {
		dto.ErrorResponse(ctx, err)
		return
	}

	dto.SuccessResponse(ctx, "Logout successfully", nil)
}

func (ah *AuthHandler) RefreshToken(ctx *gin.Context) {
	var params dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	accessToken, refreshToken, err := ah.service.RefreshToken(ctx.Request.Context(), params.RefreshToken)
	if err != nil {
		dto.ErrorResponse(ctx, err)
		return
	}

	response := dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	dto.SuccessResponse(ctx, "Refresh token successfully", response)
}

func (ah *AuthHandler) ForgotPassword(ctx *gin.Context) {
	var params dto.ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	err := ah.service.ForgotPassword(ctx.Request.Context(), params.Email)
	if err != nil {
		dto.ErrorResponse(ctx, err)
		return
	}

	dto.SuccessResponse(ctx, "Reset link sent to your email", true)
}

func (ah *AuthHandler) ResetPassword(ctx *gin.Context) {
	var params dto.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	err := ah.service.ResetPassword(ctx.Request.Context(), params.Token, params.Password)
	if err != nil {
		dto.ErrorResponse(ctx, err)
		return
	}

	dto.SuccessResponse(ctx, "Password reset successfully", true)
}
