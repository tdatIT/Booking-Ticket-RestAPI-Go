package user

import (
	"Booking-Ticket-App/domain/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"strings"
	"time"
)

type UserUsecase struct {
	userRepo       model.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(mr model.UserRepository, timeout time.Duration) model.UserUsecase {
	return &UserUsecase{
		userRepo:       mr,
		contextTimeout: timeout,
	}
}

func (u *UserUsecase) FindUserById(id string, ctx context.Context) (*model.Users, error) {
	return u.userRepo.FindUserById(id, ctx)
}

func (u *UserUsecase) FindUserByEmail(email string, ctx context.Context) (*model.Users, error) {
	return u.userRepo.FindUserByEmail(email, ctx)
}

func (u *UserUsecase) SignUp(body model.SignupBody, ctx context.Context) (*model.Users, error) {
	var user model.Users
	user.Email = strings.ToLower(body.Email)
	hashedPassword, _ := model.HashPassword(body.Password)
	user.Password = hashedPassword
	user.Address = body.Address
	user.FullName = body.FullName
	//set user status
	user.ID = primitive.NewObjectID()
	user.CreateAt = time.Now()
	user.UpdateAt = user.CreateAt
	user.IsActive = true
	user.Role = append(user.Role, "user")

	newUser, err := u.userRepo.InsertNewUser(user, ctx)

	if err != nil {
		log.Print(fmt.Errorf("not found user cause: %w ", err))
		return nil, err

	}
	return newUser, nil
}

func (u *UserUsecase) SignIn(email string, password string, ctx context.Context) (*model.Users, error) {
	user, err := u.userRepo.FindUserByEmail(email, ctx)
	if err != nil {
		log.Print(fmt.Errorf("not found user : %w ", err))
		return nil, err
	}
	if err = model.VerifyPassword(user.Password, password); err != nil {
		log.Print(fmt.Errorf("login failed check username and password:%w", err))
		return nil, err
	}
	return user, nil
}
