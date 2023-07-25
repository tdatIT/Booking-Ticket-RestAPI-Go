package routes

import (
	"Booking-Ticket-App/controller"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	authController controller.AuthController
}

func NewAuthRouter(authController controller.AuthController) AuthRouter {
	return AuthRouter{authController}
}
func (r *AuthRouter) AuthRouteRegister(rt *gin.Engine) {
	router := rt.Group("/api/v1/auth")
	router.POST("/signin", r.authController.SignInUser)
	router.POST("/signup", r.authController.SignUpUser)
}
