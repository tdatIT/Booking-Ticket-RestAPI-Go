package routes

import (
	"Booking-Ticket-App/controller"
	"Booking-Ticket-App/middleware"
	"Booking-Ticket-App/services"
	"github.com/gin-gonic/gin"
)

type MovieRoute struct {
	movieController controller.MovieController
}

func NewMovieRouter(postController controller.MovieController) MovieRoute {
	return MovieRoute{postController}
}

func (r *MovieRoute) MovieRouteRegister(rt *gin.Engine, service services.UserService) {

	router := rt.Group("/api/v1/movies")
	router.GET("/", r.movieController.GetAllMovies)
	router.GET("/:id", r.movieController.FindPostById)
	router.POST("/", middleware.O2Auth(service), r.movieController.CreateNewMovie)
	router.PATCH("/:id", middleware.O2Auth(service), r.movieController.UpdateMovie)
	router.DELETE("/:id", middleware.O2Auth(service), r.movieController.DeleteMovies)
	router.GET("/search", r.movieController.SearchMovies)
}
