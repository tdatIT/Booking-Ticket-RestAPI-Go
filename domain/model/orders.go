package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID   `bson:"user_id,omitempty" json:"user_id"`
	Total     int64                `bson:"total,omitempty" json:"total"`
	Status    string               `bson:"status,omitempty" json:"status"`
	CreatedAt time.Time            `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time            `bson:"updated_at,omitempty" json:"updated_at"`
	Tickets   []primitive.ObjectID `bson:"tickets,omitempty" json:"tickets"`
}
type OrderDTO struct {
	ID              string   `json:"id"`
	UserID          string   `json:"user_id"`
	Total           int64    `json:"total"`
	MovieScheduleID string   `json:"movie_schedule_id"`
	TicketSeats     []string `json:"tickets"`
	TicketPrice     int      `json:"ticket_price"`
}
type OrderUsecase interface {
	CreateNewOrder(dto *OrderDTO) error
}
type OrderDeliveryMessage interface {
	SendMessageToConsumer(dto *OrderDTO) error
}
