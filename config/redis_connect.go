package config

import (
	"Booking-Ticket-App/domain/model"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

func GetConnectToRedis() *redis.Client {
	var cache = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return cache
}

func ParseRedisObj(val []byte) *model.Movies {
	var movie *model.Movies
	err := json.Unmarshal(val, &movie)
	if err != nil {
		panic(err)
	}
	return movie
}
func ParseRedisListObj(val []byte) []model.Movies {
	var movies []model.Movies
	err := json.Unmarshal(val, &movies)
	if err != nil {
		panic(err)
	}
	return movies
}
