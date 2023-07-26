package controller

import (
	"Booking-Ticket-App/data/model"
	"Booking-Ticket-App/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type MovieScheduleController struct {
	movieScheduleService services.MovieScheduleService
}

func NewMovieScheduleController(mSServices services.MovieScheduleService) MovieScheduleController {
	return MovieScheduleController{movieScheduleService: mSServices}
}

func (mc *MovieScheduleController) CreateNewSM(ctx *gin.Context) {
	var schedule *model.MovieSchedule
	if err := ctx.ShouldBindJSON(&schedule); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err := mc.movieScheduleService.InsertSchedule(schedule)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.String(http.StatusOK, "new movie schedule was created ")

}

func (mc *MovieScheduleController) GetAllScheduleByMovieId(ctx *gin.Context) {
	movieId := ctx.Param("id")
	movies, err := mc.movieScheduleService.GetAllScheduleByMovie(movieId)
	if err != nil {
		ctx.String(http.StatusNotFound, "Empty schedule")
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

func (mc *MovieScheduleController) UpdateSchedule(ctx *gin.Context) {
	movieId := ctx.Param("id")
	var schedule model.MovieSchedule
	if err := ctx.ShouldBindJSON(&schedule); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err := mc.movieScheduleService.UpdateSchedule(movieId, schedule)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.String(http.StatusOK, "movie schedule was updated")
}

func (mc *MovieScheduleController) CancelSchedule(ctx *gin.Context) {
	movieId := ctx.Param("id")

	err := mc.movieScheduleService.CancelScheduleMovie(movieId)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.String(http.StatusOK, "movie schedule was deleted")
}

func (mc *MovieScheduleController) FindPostById(ctx *gin.Context) {
	scheduleId := ctx.Param("id")
	movie, err := mc.movieScheduleService.FindScheduleMovieById(scheduleId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": movie})
}
