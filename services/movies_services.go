package services

import (
	"Booking-Ticket-App/config"
	"Booking-Ticket-App/data/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type MovieServices struct {
	dbCollection *mongo.Collection
	client       *mongo.Client
	ctx          context.Context
}

func NewMovieClient(client *mongo.Client, ctx context.Context) MovieServices {
	return MovieServices{
		client:       client,
		ctx:          ctx,
		dbCollection: config.GetCollection(client, "movies"),
	}
}
func (m *MovieServices) InsertMovie(movie *model.Movies) error {
	movie.ID = primitive.NewObjectID()
	movie.CreateAt = time.Now()
	_, err := m.dbCollection.InsertOne(m.ctx, movie)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}
func (m *MovieServices) GetAllMovie() ([]model.Movies, error) {
	movies := make([]model.Movies, 0)
	rs, err := m.dbCollection.Find(m.ctx, bson.M{})
	if err != nil {
		log.Print(fmt.Errorf("could not get all movies: %w", err))
		return nil, err
	}
	if err = rs.All(m.ctx, &movies); err != nil {
		log.Print(fmt.Errorf("could marshall the movies results: %w", err))
		return nil, err
	}
	return movies, err
}
func (m *MovieServices) FindById(id string) (*model.Movies, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": objId}
	var movie *model.Movies
	err := m.dbCollection.FindOne(m.ctx, query).Decode(&movie)
	if err != nil {
		log.Print(fmt.Errorf("could not get all movies: %w", err))
		return nil, err
	}
	return movie, err
}
func (m *MovieServices) UpdateMovie(id string, movies model.Movies) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := m.dbCollection.UpdateOne(m.ctx, bson.M{"_id": objID}, bson.D{{ //nolint:govet
		"$set", bson.D{
			{"name", movies.Name},
			{"description", movies.Description},
			{"duration", movies.Duration},
			{"update_at", movies.UpdateAt},
			{"casts", movies.Casts},
			{"genres", movies.Genres},
		},
	}})

	if err != nil {
		log.Print(fmt.Errorf("could not movies book with id [%s]: %w", id, err))
		return err
	}
	return nil
}
func (m *MovieServices) DeleteMovie(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := m.dbCollection.DeleteOne(m.ctx, bson.M{"_id": objID})
	if err != nil {
		log.Print(fmt.Errorf("could not delete movies book with id [%s]: %w", id, err))
		return err
	}
	return nil
}
func (m *MovieServices) FindMovieByKeyword(keyword string) ([]model.Movies, error) {
	movies := make([]model.Movies, 0)
	rs, err := m.dbCollection.Find(m.ctx, bson.D{{"$text", bson.D{{"$search", keyword}}}})
	if err != nil {
		log.Print(fmt.Errorf("could not get all movies: %w", err))
		return nil, err
	}
	if err = rs.All(m.ctx, &movies); err != nil {
		log.Print(fmt.Errorf("could marshall the movies results: %w", err))
		return nil, err
	}
	return movies, nil
}
