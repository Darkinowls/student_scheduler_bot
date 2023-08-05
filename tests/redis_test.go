package tests

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"log"
	"studentBot/features/telegram_bot/use_case"
	"testing"
)

func TestRedisWithKeys(t *testing.T) {

	env := "../.env"
	client := getRedis(&env)
	defer client.Close()

	ctx := context.Background()

	client.FlushAll(ctx).Result()

	hashMap := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	// Start a Redis pipeline
	pipe := client.Pipeline()

	// Save the hashmap in Redis using the pipeline
	for key, value := range hashMap {
		pipe.Set(ctx, key, value, 0)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Println("Error saving hashmap in Redis:", err)
		return
	}

	log.Println("Hashmap saved in Redis successfully.")

	var cursor uint64

	keys, _, _ := client.Scan(ctx, cursor, "*", 10).Result()

	log.Println(keys)

}

func getRedis(env *string) *redis.Client {
	if env == nil {
		*env = ".env"
	}
	_ = godotenv.Overload("../.env")
	redisUrl := use_case.GetEnv("REDIS_URL")
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opt)
}
