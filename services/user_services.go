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
)

type UserService struct {
	client         *mongo.Client
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(ctx context.Context, client *mongo.Client) UserService {
	return UserService{
		client:         client,
		userCollection: config.GetCollection(client, "users"),
		ctx:            ctx,
	}
}

func (us *UserService) FindUserById(id string) (*model.Users, error) {
	ObjID, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": ObjID}
	var user *model.Users
	err := us.userCollection.FindOne(us.ctx, query).Decode(&user)
	if err != nil {
		log.Print(fmt.Errorf("not found user cause: %w ", err))
		return nil, err
	}
	return user, nil
}

func (us *UserService) FindUserByEmail(email string) (*model.Users, error) {
	query := bson.M{"email": email}
	var user *model.Users
	err := us.userCollection.FindOne(us.ctx, query).Decode(&user)
	if err != nil {
		log.Print(fmt.Errorf("not found user cause: %w ", err))
		return nil, err
	}
	return user, nil
}
