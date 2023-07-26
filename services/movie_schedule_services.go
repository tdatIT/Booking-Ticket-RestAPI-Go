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

type MovieScheduleService struct {
	dbCollection *mongo.Collection
	client       *mongo.Client
	ctx          context.Context
}

func NewMovieScheduleService(client *mongo.Client, ctx context.Context) MovieScheduleService {
	return MovieScheduleService{
		client:       client,
		ctx:          ctx,
		dbCollection: config.GetCollection(client, "schedule_movies"),
	}
}
func (ms *MovieScheduleService) InsertSchedule(sm *model.MovieSchedule) error {
	sm.ID = primitive.NewObjectID()
	sm.CreatedAt = time.Now()
	_, err := ms.dbCollection.InsertOne(ms.ctx, sm)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}
func (ms *MovieScheduleService) GetAllScheduleByMovie(Id string) ([]model.MovieSchedule, error) {
	schedules := make([]model.MovieSchedule, 0)
	movieId, _ := primitive.ObjectIDFromHex(Id)
	query := bson.M{"movie_id": movieId}
	rs, err := ms.dbCollection.Find(ms.ctx, query)
	if err != nil {
		log.Print(fmt.Errorf("Not found schedule for movie cause: %w ", err))
		return nil, err
	}
	if err = rs.All(ms.ctx, &schedules); err != nil {
		log.Print(fmt.Errorf("could marshall the schedule results: %w", err))
		return nil, err
	}
	return schedules, err
}

func (ms *MovieScheduleService) FindScheduleMovieById(id string) (*model.MovieSchedule, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": objId}
	var schedule *model.MovieSchedule
	err := ms.dbCollection.FindOne(ms.ctx, query).Decode(&schedule)
	if err != nil {
		log.Print(fmt.Errorf("not found schedule cause: %w", err))
		return nil, err
	}
	return schedule, err
}
func (ms *MovieScheduleService) UpdateSchedule(id string, sm model.MovieSchedule) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := ms.dbCollection.UpdateOne(ms.ctx, bson.M{"_id": objID}, bson.D{{
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

func (ms *MovieScheduleService) CancelScheduleMovie(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := ms.dbCollection.UpdateOne(ms.ctx, bson.M{"_id": objID}, bson.D{{
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
