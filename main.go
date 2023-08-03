package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"os"
	httpserver "studentBot/features/http_server/delivery"
	telegrambot "studentBot/features/telegram_bot/delivery"
	"studentBot/features/telegram_bot/models"
	"time"
)

func main() {

	_ = godotenv.Overload()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	botKey := os.Getenv("BOT_KEY")
	if botKey == "" {
		panic("No BOT_KEY")
	}
	bot, err := tgbotapi.NewBotAPI(botKey)
	if err != nil {
		panic(err)
	}

	go httpserver.ServeHttpServer(&port)

	const MyChatId int64 = 1647688266
	scheduleMap := make(map[string]*models.ScheduleEntity)
	go telegrambot.CheckUpdates(bot, MyChatId, scheduleMap)
	telegrambot.RunPeriodically(2*time.Minute, bot, scheduleMap, MyChatId)

}
