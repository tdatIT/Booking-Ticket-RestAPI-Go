package services

import (
	"Booking-Ticket-App/config"
	"Booking-Ticket-App/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"time"
)

type AuthService struct {
	client         *mongo.Client
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewAuthService(client *mongo.Client, ctx context.Context) AuthService {
	return AuthService{
		client:         client,
		userCollection: config.GetCollection(client, "users"),
		ctx:            ctx,
	}
}

func (as *AuthService) SignUp(body model.SignupBody) (*model.Users, error) {
	var user model.Users
	user.Email = strings.ToLower(body.Email)
	hashedPassword, _ := HashPassword(body.Password)
	user.Password = hashedPassword
	user.Address = body.Address
	user.FullName = body.FullName
	//set user status
	user.ID = primitive.NewObjectID()
	user.CreateAt = time.Now()
	user.UpdateAt = user.CreateAt
	user.IsActive = true
	user.Role = append(user.Role, "user")

	res, err := as.userCollection.InsertOne(as.ctx, user)

	if err != nil {
		log.Print(fmt.Errorf("not found user cause: %w ", err))
		return nil, err

	}
	var newUser *model.Users
	query := bson.M{"_id": res.InsertedID}

	err = as.userCollection.FindOne(as.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
func (as *AuthService) SignIn(email string, password string) (*model.Users, error) {
	query := bson.M{"email": email}
	var user *model.Users
	err := as.userCollection.FindOne(as.ctx, query).Decode(&user)
	if err != nil {
		log.Print(fmt.Errorf("not found user : %w ", err))
		return nil, err
	}
	if err = VerifyPassword(user.Password, password); err != nil {
		log.Print(fmt.Errorf("login failed check username and password:%w", err))
		return nil, err
	}
	return user, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}
func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
