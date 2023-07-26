package controller

import (
	"Booking-Ticket-App/config"
	model2 "Booking-Ticket-App/data/model"
	"Booking-Ticket-App/helper"
	"Booking-Ticket-App/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthController struct {
	authService services.AuthService
	userService services.UserService
}

func NewAuthController(authService services.AuthService, userService services.UserService) AuthController {
	return AuthController{authService, userService}
}

func (ac *AuthController) SignInUser(ctx *gin.Context) {
	var loginBody *model2.LoginDTO

	if err := ctx.ShouldBindJSON(&loginBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	user, err := ac.authService.SignIn(loginBody.Email, loginBody.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	// Generate Tokens
	accessToken, err := helper.CreateToken(user.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}

func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var body model2.SignupBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if body.Password != body.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	newUser, err := ac.authService.SignUp(body)

	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": newUser}})
}
