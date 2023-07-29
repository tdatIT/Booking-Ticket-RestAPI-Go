package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MovieSchedule struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	MovieID   primitive.ObjectID `bson:"movie_id,omitempty" json:"movie_id"`
	StartTime time.Time          `bson:"start_time,omitempty" json:"start_time"`
	EndTime   time.Time          `bson:"end_time,omitempty" json:"end_time"`
	Room      string             `bson:"room,omitempty" json:"room"`
	Price     int64              `bson:"price,omitempty" json:"price"`
	Status    string             `bson:"status,omitempty" json:"status"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
	IsCancel  bool               `bson:"is_cancel,omitempty" json:"is_cancel"`
}

type MovieScheduleRepository interface {
	InsertSchedule(sm *MovieSchedule, ctx context.Context) error
	GetAllScheduleByMovie(Id string, ctx context.Context) ([]MovieSchedule, error)
	FindScheduleMovieById(id string, ctx context.Context) (*MovieSchedule, error)
	UpdateSchedule(id string, sm MovieSchedule, ctx context.Context) error
	CancelScheduleMovie(id string, ctx context.Context) error
}

type MovieScheduleUsecase interface {
	InsertSchedule(sm *MovieSchedule, ctx context.Context) error
	GetAllScheduleByMovie(Id string, ctx context.Context) ([]MovieSchedule, error)
	FindScheduleMovieById(id string, ctx context.Context) (*MovieSchedule, error)
	UpdateSchedule(id string, sm MovieSchedule, ctx context.Context) error
	CancelScheduleMovie(id string, ctx context.Context) error
}
