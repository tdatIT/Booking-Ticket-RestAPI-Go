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
)

type MongoDBUserRepository struct {
	client         *mongo.Client
	userCollection *mongo.Collection
}

func NewMongoDBUserRepository(client *mongo.Client) model.UserRepository {
	return &MongoDBUserRepository{
		client:         client,
		userCollection: config.GetCollection(client, "users"),
	}
}

func (us *MongoDBUserRepository) FindUserById(id string, ctx context.Context) (*model.Users, error) {
	ObjID, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": ObjID}
	var user *model.Users
	err := us.userCollection.FindOne(ctx, query).Decode(&user)
	if err != nil {
		log.Print(fmt.Errorf("not found user cause: %w ", err))
		return nil, err
	}
	return user, nil
}

func (us *MongoDBUserRepository) FindUserByEmail(email string, ctx context.Context) (*model.Users, error) {
	query := bson.M{"email": email}
	var user *model.Users
	err := us.userCollection.FindOne(ctx, query).Decode(&user)
	if err != nil {
		log.Print(fmt.Errorf("not found user cause: %w ", err))
		return nil, err
	}
	return user, nil
}

func (us *MongoDBUserRepository) InsertNewUser(user model.Users, ctx context.Context) (*model.Users, error) {
	res, err := us.userCollection.InsertOne(ctx, user)
	if err != nil {
		log.Print(fmt.Errorf("not found user cause: %w ", err))
		return nil, err

	}
	var newUser *model.Users
	query := bson.M{"_id": res.InsertedID}
	err = us.userCollection.FindOne(ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
