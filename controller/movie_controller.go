package controller

import (
	"Booking-Ticket-App/data/model"
	"Booking-Ticket-App/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type MovieController struct {
	movieServices services.MovieServices
}

func NewMovieController(movieServices services.MovieServices) MovieController {
	return MovieController{movieServices: movieServices}
}

func (mc *MovieController) CreateNewMovie(ctx *gin.Context) {
	var movie *model.Movies
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err := mc.movieServices.InsertMovie(movie)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.String(http.StatusOK, "Create new movie was success")

}

func (mc *MovieController) GetAllMovies(ctx *gin.Context) {
	movies, err := mc.movieServices.GetAllMovie()
	if err != nil {
		ctx.String(http.StatusNotFound, "Empty movie")
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

func (mc *MovieController) SearchMovies(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	movies, err := mc.movieServices.FindMovieByKeyword(keyword)
	if err != nil {
		ctx.String(http.StatusNotFound, "Empty movie")
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

func (mc *MovieController) UpdateMovie(ctx *gin.Context) {
	movieId := ctx.Param("id")
	var movie model.Movies
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err := mc.movieServices.UpdateMovie(movieId, movie)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.String(http.StatusOK, "movie was updated")
}
func (mc *MovieController) DeleteMovies(ctx *gin.Context) {
	movieId := ctx.Param("id")

	err := mc.movieServices.DeleteMovie(movieId)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.String(http.StatusOK, "Movie was deleted")
}
func (mc *MovieController) FindPostById(ctx *gin.Context) {
	movieId := ctx.Param("id")
	movie, err := mc.movieServices.FindById(movieId)
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
