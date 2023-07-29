package movie

import (
	"Booking-Ticket-App/domain/model"
	"context"
	"time"
)

type MovieUsecase struct {
	movieRepo      model.MovieRepository
	cachingRepo    model.CachingMovieRepository
	contextTimeout time.Duration
}

func NewMovieUseCase(mr model.MovieRepository, cachingRepo model.CachingMovieRepository, timeout time.Duration) model.MovieUsecase {
	return &MovieUsecase{
		movieRepo:      mr,
		cachingRepo:    cachingRepo,
		contextTimeout: timeout,
	}
}

func (m *MovieUsecase) InsertMovie(movie *model.Movies, c context.Context) (*model.Movies, error) {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()

	new_movie, err := m.movieRepo.InsertMovie(movie, ctx)
	if err != nil {
		return nil, err
	}

	err = m.cachingRepo.SetKeyValue(movie, ctx)
	if err != nil {
		return nil, err
	}
	return new_movie, nil
}

func (m *MovieUsecase) GetAllMovie(c context.Context) ([]model.Movies, error) {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()
	list, err := m.movieRepo.GetAllMovie(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (m *MovieUsecase) FindById(id string, c context.Context) (*model.Movies, error) {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()
	movie, err := m.cachingRepo.GetByKey(id, ctx)
	if err != nil {
		movie, err = m.movieRepo.FindById(id, ctx)
		if err != nil {
			return nil, err
		}
		m.cachingRepo.SetKeyValue(movie, ctx)
	}
	return movie, nil
}

func (m *MovieUsecase) UpdateMovie(id string, movies *model.Movies, c context.Context) error {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()
	m.cachingRepo.SetKeyValue(movies, ctx)
	return m.movieRepo.UpdateMovie(id, movies, ctx)
}

func (m *MovieUsecase) DeleteMovie(id string, c context.Context) error {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()
	m.cachingRepo.ClearCache(id, ctx)
	return m.movieRepo.DeleteMovie(id, ctx)
}

func (m *MovieUsecase) FindMovieByKeyword(keyword string, c context.Context) ([]model.Movies, error) {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()
	list, err := m.cachingRepo.GetListByKey(keyword, ctx)
	if err != nil {
		list, err = m.movieRepo.FindMovieByKeyword(keyword, ctx)
		if err != nil {
			return nil, err
		}
		m.cachingRepo.SetListValue(keyword, list, ctx)
	}
	return list, nil
}
