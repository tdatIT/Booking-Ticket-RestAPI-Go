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
