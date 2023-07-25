package helper

import (
	"Booking-Ticket-App/config"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

type SignedDetails struct {
	Email string
	jwt.StandardClaims
}

func CreateToken(Email string) (signedToken string, err error) {
	claims := &SignedDetails{
		Email: Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.SecretKey))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, err
}

// ValidateToken validates the jwt token
func ValidateToken(signedToken string) (claims *SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.SecretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		log.Print("the token is invalid")
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		log.Print("token is expired")
		return nil, err
	}
	return claims, err
}
