package controllers

import (
	"github.com/Tretorhate/university-management-system/internal/dto"
	"github.com/Tretorhate/university-management-system/internal/service"
	"github.com/Tretorhate/university-management-system/pkg/errors"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var request dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(errors.BadRequest("Invalid request body", err))
		return
	}

	response, err := c.authService.Register(&request)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(201, response)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var request dto.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(errors.BadRequest("Invalid request body", err))
		return
	}

	response, err := c.authService.Login(&request)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, response)
}
