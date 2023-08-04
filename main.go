package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"strconv"
	httpserver "studentBot/features/http_server/delivery"
	"studentBot/features/telegram_bot/delivery"
	"studentBot/features/telegram_bot/models"
)

func main() {

	_ = godotenv.Overload()

	port := delivery.GetEnv("PORT")

	botKey := delivery.GetEnv("BOT_KEY")

	MyChatId, err := strconv.ParseInt(delivery.GetEnv("CHAT_ID"), 10, 64)
	if err != nil {
		panic("CHAT_ID error:" + err.Error())
	}

	redisUrl := delivery.GetEnv("REDIS_URL")
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(botKey)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)
	defer client.Close()

	go httpserver.ServeHttpServer(&port)

	scheduleMap := make(map[string]*models.ScheduleEntity)
	go delivery.CheckUpdates(bot, MyChatId, scheduleMap)
	delivery.RunPeriodically(bot, scheduleMap)

}
