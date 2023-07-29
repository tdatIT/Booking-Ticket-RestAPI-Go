package repository

import (
	"Booking-Ticket-App/config"
	"Booking-Ticket-App/domain/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type MongodbMovieRepository struct {
	dbCollection *mongo.Collection
	client       *mongo.Client
}

func NewMongodbMovieRepository(client *mongo.Client) MongodbMovieRepository {
	return MongodbMovieRepository{
		client:       client,
		dbCollection: config.GetCollection(client, "movies"),
	}
}

func (m MongodbMovieRepository) InsertMovie(movie *model.Movies, ctx context.Context) (*model.Movies, error) {
	movie.ID = primitive.NewObjectID()
	movie.CreateAt = time.Now()
	newId, err := m.dbCollection.InsertOne(ctx, movie)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	query := bson.M{"_id": newId.InsertedID}
	var new_movie *model.Movies
	err = m.dbCollection.FindOne(ctx, query).Decode(&new_movie)
	if err != nil {
		return nil, err
	}
	return new_movie, nil

}

func (m MongodbMovieRepository) GetAllMovie(ctx context.Context) ([]model.Movies, error) {
	movies := make([]model.Movies, 0)
	rs, err := m.dbCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Print(fmt.Errorf("could not get all movies: %w", err))
		return nil, err
	}
	if err = rs.All(ctx, &movies); err != nil {
		log.Print(fmt.Errorf("could marshall the movies results: %w", err))
		return nil, err
	}
	return movies, err
}

func (m MongodbMovieRepository) FindById(id string, ctx context.Context) (*model.Movies, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": objId}
	var movie *model.Movies
	err := m.dbCollection.FindOne(ctx, query).Decode(&movie)
	if err != nil {
		log.Print(fmt.Errorf("could not get all movies: %w", err))
		return nil, err
	}
	return movie, err
}

func (m MongodbMovieRepository) UpdateMovie(id string, movies *model.Movies, ctx context.Context) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := m.dbCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.D{{ //nolint:govet
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

func (m MongodbMovieRepository) DeleteMovie(id string, ctx context.Context) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := m.dbCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		log.Print(fmt.Errorf("could not delete movies book with id [%s]: %w", id, err))
		return err
	}
	return nil
}

func (m MongodbMovieRepository) FindMovieByKeyword(keyword string, ctx context.Context) ([]model.Movies, error) {
	movies := make([]model.Movies, 0)
	rs, err := m.dbCollection.Find(ctx, bson.D{{"$text", bson.D{{"$search", keyword}}}})
	if err != nil {
		log.Print(fmt.Errorf("could not get all movies: %w", err))
		return nil, err
	}
	if err = rs.All(ctx, &movies); err != nil {
		log.Print(fmt.Errorf("could marshall the movies results: %w", err))
		return nil, err
	}
	return movies, nil
}
