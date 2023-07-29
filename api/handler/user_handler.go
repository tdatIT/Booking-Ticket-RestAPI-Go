package handler

import (
	"Booking-Ticket-App/config"
	model2 "Booking-Ticket-App/domain/model"
	"Booking-Ticket-App/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type UserHandler struct {
	userUsecase model2.UserUsecase
}

func NewUserHandler(userUsecase model2.UserUsecase, ctx *gin.Engine) {
	handler := UserHandler{
		userUsecase: userUsecase,
	}
	route := ctx.Group("/api/v1/auth")
	route.POST("/signin", handler.SignInUser)
	route.POST("/signup", handler.SignUpUser)
}

func (ac *UserHandler) SignInUser(ctx *gin.Context) {
	var loginBody *model2.LoginDTO

	if err := ctx.ShouldBindJSON(&loginBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	user, err := ac.userUsecase.SignIn(loginBody.Email, loginBody.Password, ctx.Request.Context())
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

func (ac *UserHandler) SignUpUser(ctx *gin.Context) {
	var body model2.SignupBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if body.Password != body.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	newUser, err := ac.userUsecase.SignUp(body, ctx.Request.Context())

	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "domain": gin.H{"user": newUser}})
}
