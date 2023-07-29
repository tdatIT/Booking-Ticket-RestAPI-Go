package repository

import (
	"Booking-Ticket-App/config"
	"Booking-Ticket-App/domain/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type RedisCachingData struct {
	redisClient *redis.Client
}

func NewRedisCachingRepo(client *redis.Client) RedisCachingData {
	return RedisCachingData{redisClient: client}
}
func (k RedisCachingData) SetKeyValue(movie *model.Movies, ctx context.Context) error {
	data, _ := json.Marshal(movie)
	err := k.redisClient.Set(ctx, movie.ID.String(), data, 120*time.Second).Err()
	if err != nil {
		log.Print(fmt.Errorf("cannot set movie to cache cause:%w", err))
		return err
	}
	log.Printf("Set data to cache - key:%v", movie.ID.String())
	return nil
}

func (k RedisCachingData) GetByKey(movieId string, ctx context.Context) (*model.Movies, error) {
	movie, err := k.redisClient.Get(ctx, fmt.Sprintf("ObjectID(%q)", movieId)).Bytes()
	if err != nil {
		log.Print(fmt.Errorf("cannot get movie to cache cause:%w", err))
		return nil, err
	}
	data := config.ParseRedisObj(movie)
	log.Printf("Fetch data from cache - key:%v", movieId)
	return data, nil
}

func (k RedisCachingData) ClearCache(movieId string, ctx context.Context) error {
	err := k.redisClient.Del(ctx, fmt.Sprintf("ObjectID(%q)", movieId)).Err()
	if err != nil {
		log.Print(fmt.Errorf("cannot get movie to cache cause:%w", err))
		return err
	}
	log.Printf("Clear data from cache - key:%v", movieId)
	return nil
}
func (k RedisCachingData) SetListValue(key string, movie []model.Movies, ctx context.Context) error {
	data, _ := json.Marshal(movie)
	err := k.redisClient.Set(ctx, key, data, 120*time.Second).Err()
	if err != nil {
		log.Print(fmt.Errorf("cannot set movie to cache cause:%w", err))
		return err
	}
	log.Printf("Set data to cache - key:%v", key)
	return nil
}
func (k RedisCachingData) GetListByKey(key string, ctx context.Context) ([]model.Movies, error) {
	list, err := k.redisClient.Get(ctx, fmt.Sprintf(key)).Bytes()
	if err != nil {
		log.Print(fmt.Errorf("cannot get movie to cache cause:%w", err))
		return nil, err
	}
	data := config.ParseRedisListObj(list)
	log.Printf("Fetch data from cache - key:%v", key)
	return data, nil
}
