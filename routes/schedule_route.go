package routes

import (
	"Booking-Ticket-App/controller"
	"Booking-Ticket-App/middleware"
	"Booking-Ticket-App/services"
	"github.com/gin-gonic/gin"
)

type MovieScheduleRoute struct {
	movieScheduleController controller.MovieScheduleController
}

func NewMovieScheduleRoute(movieScheduleController controller.MovieScheduleController) MovieScheduleRoute {
	return MovieScheduleRoute{movieScheduleController}
}

func (r *MovieScheduleRoute) MovieScheduleRouteRegister(rt *gin.Engine, service services.UserService) {

	router := rt.Group("/api/v1/movie_schedule")
	router.GET("/movie/:id", r.movieScheduleController.GetAllScheduleByMovieId)
	router.GET("/:id", r.movieScheduleController.FindPostById)
	//Admin
	router.POST("/", middleware.O2Auth(service), r.movieScheduleController.CreateNewSM)
	router.PATCH("/:id", middleware.O2Auth(service), r.movieScheduleController.UpdateSchedule)
	router.DELETE("/:id", middleware.O2Auth(service), r.movieScheduleController.CancelSchedule)
}
