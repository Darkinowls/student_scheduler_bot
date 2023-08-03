package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	httpserver "studentBot/features/http_server/delivery"
	telegrambot "studentBot/features/telegram_bot/delivery"
	"studentBot/features/telegram_bot/models"
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
	MyChatId, err := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)
	if err != nil {
		panic("CHAT_ID error:" + err.Error())
	}
	bot, err := tgbotapi.NewBotAPI(botKey)
	if err != nil {
		panic(err)
	}

	go httpserver.ServeHttpServer(&port)

	scheduleMap := make(map[string]*models.ScheduleEntity)
	go telegrambot.CheckUpdates(bot, MyChatId, scheduleMap)
	telegrambot.RunPeriodically(bot, scheduleMap)

}
