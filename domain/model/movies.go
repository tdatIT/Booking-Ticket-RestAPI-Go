package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Movies struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Director    string             `bson:"director,omitempty" json:"director"`
	Casts       []string           `bson:"casts,omitempty" json:"casts"`
	Genres      []string           `bson:"genres,omitempty" json:"genres"`
	Duration    int                `bson:"duration,omitempty" json:"duration"`
	ReleaseDate time.Time          `bson:"release_date,omitempty" json:"release_date"`
	Description string             `bson:"description,omitempty" json:"description"`
	ImagesUrl   []string           `bson:"images_url,omitempty" json:"images_url"`
	CreateAt    time.Time          `bson:"create_at,omitempty" json:"create_at"`
	UpdateAt    time.Time          `bson:"update_at,omitempty" json:"update_at"`
}

type MovieUsecase interface {
	InsertMovie(movie *Movies, c context.Context) (*Movies, error)
	GetAllMovie(c context.Context) ([]Movies, error)
	FindById(id string, c context.Context) (*Movies, error)
	UpdateMovie(id string, movies *Movies, c context.Context) error
	DeleteMovie(id string, c context.Context) error
	FindMovieByKeyword(keyword string, c context.Context) ([]Movies, error)
}

type MovieRepository interface {
	InsertMovie(movie *Movies, ctx context.Context) (*Movies, error)
	GetAllMovie(ctx context.Context) ([]Movies, error)
	FindById(id string, ctx context.Context) (*Movies, error)
	UpdateMovie(id string, movies *Movies, ctx context.Context) error
	DeleteMovie(id string, ctx context.Context) error
	FindMovieByKeyword(keyword string, ctx context.Context) ([]Movies, error)
}
type CachingMovieRepository interface {
	SetKeyValue(movie *Movies, ctx context.Context) error
	GetByKey(id string, ctx context.Context) (*Movies, error)
	GetListByKey(key string, ctx context.Context) ([]Movies, error)
	SetListValue(key string, movie []Movies, ctx context.Context) error
	ClearCache(id string, ctx context.Context) error
}
