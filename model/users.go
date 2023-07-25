package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Users struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `bson:"email,omitempty" json:"email"`
	FullName string             `bson:"full_name,omitempty" json:"full_name"`
	Password string             `bson:"password,omitempty" json:"password"`
	Address  string             `bson:"address,omitempty" json:"address"`
	Role     []string           `bson:"role,omitempty" json:"role"`
	CreateAt time.Time          `bson:"create_at,omitempty" json:"create_at"`
	UpdateAt time.Time          `bson:"update_at,omitempty" json:"update_at"`
	IsActive bool               `bson:"is_active,omitempty" json:"is_active"`
}
