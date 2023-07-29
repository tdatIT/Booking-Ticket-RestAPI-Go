package handler

import (
	"Booking-Ticket-App/api/middleware"
	"Booking-Ticket-App/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type MovieScheduleHandler struct {
	movieScheduleUsecase model.MovieScheduleUsecase
	userRepository       model.UserRepository
}

func NewMovieScheduleHandler(movieScheduleUsecase model.MovieScheduleUsecase, userRepository model.UserRepository, ctx *gin.Engine) {
	handler := MovieScheduleHandler{movieScheduleUsecase: movieScheduleUsecase, userRepository: userRepository}
	router := ctx.Group("/api/v1/movie_schedule")
	router.GET("/movie/:id", handler.GetAllScheduleByMovieId)
	router.GET("/:id", handler.FindPostById)
	//Admin
	router.POST("/", middleware.O2Auth(userRepository), handler.CreateNewSM)
	router.PATCH("/:id", middleware.O2Auth(userRepository), handler.UpdateSchedule)
	router.DELETE("/:id", middleware.O2Auth(userRepository), handler.CancelSchedule)
}

func (mc *MovieScheduleHandler) CreateNewSM(ctx *gin.Context) {
	var schedule *model.MovieSchedule
	if err := ctx.ShouldBindJSON(&schedule); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err := mc.movieScheduleUsecase.InsertSchedule(schedule, ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.String(http.StatusOK, "new movie schedule was created ")

}

func (mc *MovieScheduleHandler) GetAllScheduleByMovieId(ctx *gin.Context) {
	movieId := ctx.Param("id")
	movies, err := mc.movieScheduleUsecase.GetAllScheduleByMovie(movieId, ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusNotFound, "Empty schedule")
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

func (mc *MovieScheduleHandler) UpdateSchedule(ctx *gin.Context) {
	movieId := ctx.Param("id")
	var schedule model.MovieSchedule
	if err := ctx.ShouldBindJSON(&schedule); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err := mc.movieScheduleUsecase.UpdateSchedule(movieId, schedule, ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.String(http.StatusOK, "movie schedule was updated")
}

func (mc *MovieScheduleHandler) CancelSchedule(ctx *gin.Context) {
	movieId := ctx.Param("id")

	err := mc.movieScheduleUsecase.CancelScheduleMovie(movieId, ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.String(http.StatusOK, "movie schedule was deleted")
}

func (mc *MovieScheduleHandler) FindPostById(ctx *gin.Context) {
	scheduleId := ctx.Param("id")
	movie, err := mc.movieScheduleUsecase.FindScheduleMovieById(scheduleId, ctx.Request.Context())
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "domain": movie})
}
