package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"studentBot/features/telegram_bot/models"
)

type ScheduleRepository interface {
	DeleteAllRecords() error
	SaveScheduleEntities(sMap map[string]*models.ScheduleEntity) error
	GetScheduleEntities() (map[string]*models.ScheduleEntity, error)
	Close() error
}

type RedisScheduleRepository struct {
	redis       *redis.Client
	scheduleMap map[string]*models.ScheduleEntity
}

func (r *RedisScheduleRepository) DeleteAllRecords() error {
	_, err := r.redis.FlushAll(context.Background()).Result()
	return err
}

func (r *RedisScheduleRepository) SaveScheduleEntities(sMap map[string]*models.ScheduleEntity) error {
	ctx := context.Background()
	pipe := r.redis.Pipeline()
	for key, value := range sMap {
		pipe.Set(ctx, key, value, 0)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Println("Error saving hashmap in Redis:", err)
		return err
	}
	return nil
}

func (r *RedisScheduleRepository) GetScheduleEntities() (map[string]*models.ScheduleEntity, error) {
	ctx := context.Background()

	// Fetch all keys and their values from Redis
	keys, err := r.redis.Keys(ctx, "*").Result()
	if err != nil {
		log.Println("Error retrieving keys:", err)
		return nil, err
	}
	// Fetch values for all keys
	scheduleMap := make(map[string]*models.ScheduleEntity)
	for _, key := range keys {
		value, err := r.redis.Get(ctx, key).Result()
		if err != nil {
			log.Printf("Error fetching value for key %s: %v\n", key, err)
			return nil, err
		}
		schedule, err := models.UnmarshalScheduleEntity([]byte(value))
		if err != nil {
			log.Printf("Error fetching value for key %s: %v\n", key, err)
			return nil, err
		}
		scheduleMap[key] = &schedule
	}

	return scheduleMap, nil
}

func (r *RedisScheduleRepository) Close() error {
	return r.redis.Close()
}

func NewScheduleRepository(redisUrl *string) ScheduleRepository {
	opt, err := redis.ParseURL(*redisUrl)
	if err != nil {
		log.Println(err)
	}
	return &RedisScheduleRepository{redis: redis.NewClient(opt)}
}
