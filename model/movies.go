package model

import (
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
