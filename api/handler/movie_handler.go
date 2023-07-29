package handler

import (
	"Booking-Ticket-App/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type MovieHandle struct {
	movieUsecase model.MovieUsecase
}

func NewMovieController(movieUsecase model.MovieRepository, ctx *gin.Engine) {
	handler := MovieHandle{
		movieUsecase: movieUsecase,
	}
	router := ctx.Group("/api/v1/movies")
	router.GET("/", handler.GetAllMovies)
	router.GET("/:id", handler.FindMovieById)
	router.POST("/", handler.CreateNewMovie)
	router.PATCH("/:id", handler.UpdateMovie)
	router.DELETE("/:id", handler.DeleteMovies)
	router.GET("/search", handler.SearchMovies)

}

func (mc *MovieHandle) FindMovieById(ctx *gin.Context) {
	movieId := ctx.Param("id")
	movie, err := mc.movieUsecase.FindById(movieId, ctx.Request.Context())
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
func (mc *MovieHandle) CreateNewMovie(ctx *gin.Context) {
	var movie *model.Movies
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	cache_movie, err := mc.movieUsecase.InsertMovie(movie, ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "\"Create new movie was success\"", "data": cache_movie.ID})
}

func (mc *MovieHandle) GetAllMovies(ctx *gin.Context) {
	movies, err := mc.movieUsecase.GetAllMovie(ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusNotFound, "Empty movie")
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

func (mc *MovieHandle) SearchMovies(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	movies, err := mc.movieUsecase.FindMovieByKeyword(keyword, ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusNotFound, "Empty movie")
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

func (mc *MovieHandle) UpdateMovie(ctx *gin.Context) {
	movieId := ctx.Param("id")
	var movie *model.Movies
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err := mc.movieUsecase.UpdateMovie(movieId, movie, ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.String(http.StatusOK, "movie was updated")
}

func (mc *MovieHandle) DeleteMovies(ctx *gin.Context) {
	movieId := ctx.Param("id")

	err := mc.movieUsecase.DeleteMovie(movieId, ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.String(http.StatusOK, "Movie was deleted")
}
