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

type MongoDBMovieScheduleRepository struct {
	dbCollection *mongo.Collection
	client       *mongo.Client
}

func NewMongoDBMovieScheduleRepository(client *mongo.Client) model.MovieScheduleRepository {
	return &MongoDBMovieScheduleRepository{
		client:       client,
		dbCollection: config.GetCollection(client, "schedule_movies"),
	}
}
func (ms *MongoDBMovieScheduleRepository) InsertSchedule(sm *model.MovieSchedule, ctx context.Context) error {
	sm.ID = primitive.NewObjectID()
	sm.CreatedAt = time.Now()
	_, err := ms.dbCollection.InsertOne(ctx, sm)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}
func (ms *MongoDBMovieScheduleRepository) GetAllScheduleByMovie(Id string, ctx context.Context) ([]model.MovieSchedule, error) {
	schedules := make([]model.MovieSchedule, 0)
	movieId, _ := primitive.ObjectIDFromHex(Id)
	query := bson.M{"movie_id": movieId}
	rs, err := ms.dbCollection.Find(ctx, query)
	if err != nil {
		log.Print(fmt.Errorf("Not found schedule for movie cause: %w ", err))
		return nil, err
	}
	if err = rs.All(ctx, &schedules); err != nil {
		log.Print(fmt.Errorf("could marshall the schedule results: %w", err))
		return nil, err
	}
	return schedules, err
}

func (ms *MongoDBMovieScheduleRepository) FindScheduleMovieById(id string, ctx context.Context) (*model.MovieSchedule, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": objId}
	var schedule *model.MovieSchedule
	err := ms.dbCollection.FindOne(ctx, query).Decode(&schedule)
	if err != nil {
		log.Print(fmt.Errorf("not found schedule cause: %w", err))
		return nil, err
	}
	return schedule, err
}
func (ms *MongoDBMovieScheduleRepository) UpdateSchedule(id string, sm model.MovieSchedule, ctx context.Context) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := ms.dbCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.D{{
		"$set", bson.D{
			{"room", sm.Room},
			{"start_time", sm.StartTime},
			{"end_time", sm.EndTime},
			{"price", sm.Price},
			{"status", sm.Status},
			{"update_at", time.Now()},
		},
	}})

	if err != nil {
		log.Print(fmt.Errorf("could not schedule movie with id [%s]: %w", id, err))
		return err
	}
	return nil
}

func (ms *MongoDBMovieScheduleRepository) CancelScheduleMovie(id string, ctx context.Context) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := ms.dbCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.D{{
		"$set", bson.D{
			{"status", "cancel"},
			{"is_cancel", true},
		},
	}})
	if err != nil {
		log.Print(fmt.Errorf("could not cancel schedule with id [%s]: %w", id, err))
		return err
	}
	return nil
}
