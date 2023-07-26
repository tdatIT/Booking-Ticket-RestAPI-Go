package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Ticket struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ShowtimeID primitive.ObjectID `bson:"showtime_id,omitempty" json:"showtime_id"`
	SeatNumber string             `bson:"seat_number,omitempty" json:"seat_number"`
	Price      int64              `bson:"price,omitempty" json:"price"`
	Status     string             `bson:"status,omitempty" json:"status"`
	CreatedAt  time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
}
