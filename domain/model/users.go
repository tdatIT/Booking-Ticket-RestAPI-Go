package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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
type LoginDTO struct {
	Email    string
	Password string
}

type SignupBody struct {
	Email           string `json:"email"`
	FullName        string `json:"fullname"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Address         string `json:"address"`
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

type UserUsecase interface {
	FindUserById(id string, ctx context.Context) (*Users, error)
	FindUserByEmail(email string, ctx context.Context) (*Users, error)
	SignUp(body SignupBody, ctx context.Context) (*Users, error)
	SignIn(email string, password string, ctx context.Context) (*Users, error)
}

type UserRepository interface {
	FindUserById(id string, ctx context.Context) (*Users, error)
	FindUserByEmail(email string, ctx context.Context) (*Users, error)
	InsertNewUser(user Users, ctx context.Context) (*Users, error)
}
