package movie_schedule

import (
	"Booking-Ticket-App/domain/model"
	"context"
	"time"
)

type MovieScheduleUsecase struct {
	movieScheduleMovie model.MovieScheduleRepository
	contextTimeout     time.Duration
}

func NewMovieScheduleUsecase(smRepository model.MovieScheduleRepository, contextTimeout time.Duration) MovieScheduleUsecase {
	return MovieScheduleUsecase{
		movieScheduleMovie: smRepository,
		contextTimeout:     contextTimeout,
	}
}

func (m MovieScheduleUsecase) InsertSchedule(sm *model.MovieSchedule, ctx context.Context) error {
	ctx_timeout, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()
	return m.movieScheduleMovie.InsertSchedule(sm, ctx_timeout)
}

func (m MovieScheduleUsecase) GetAllScheduleByMovie(Id string, ctx context.Context) ([]model.MovieSchedule, error) {
	ctx_timeout, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()
	return m.movieScheduleMovie.GetAllScheduleByMovie(Id, ctx_timeout)
}

func (m MovieScheduleUsecase) FindScheduleMovieById(id string, ctx context.Context) (*model.MovieSchedule, error) {
	ctx_timeout, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()
	return m.movieScheduleMovie.FindScheduleMovieById(id, ctx_timeout)
}

func (m MovieScheduleUsecase) UpdateSchedule(id string, sm model.MovieSchedule, ctx context.Context) error {
	ctx_timeout, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()
	return m.movieScheduleMovie.UpdateSchedule(id, sm, ctx_timeout)
}

func (m MovieScheduleUsecase) CancelScheduleMovie(id string, ctx context.Context) error {
	ctx_timeout, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()
	return m.movieScheduleMovie.CancelScheduleMovie(id, ctx_timeout)
}
