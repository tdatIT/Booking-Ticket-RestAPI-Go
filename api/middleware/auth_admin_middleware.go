package middleware

import (
	"Booking-Ticket-App/domain/model"
	"Booking-Ticket-App/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func O2Auth(userRepository model.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		}

		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		sub, err := helper.ValidateToken(access_token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		user, err := userRepository.FindUserByEmail(sub.Email, ctx.Request.Context())
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "The user belonging to this token no logger exists"})
			return
		}
		if user.Role[0] == "admin" {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "Api required role admin"})
			return
		}
	}
}
