package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"strconv"
	httpserver "studentBot/features/http_server/delivery"
	"studentBot/features/telegram_bot/consts"
	"studentBot/features/telegram_bot/delivery"
	"studentBot/features/telegram_bot/models"
	"studentBot/features/telegram_bot/use_case"
)

func main() {

	_ = godotenv.Overload()

	port := use_case.GetEnv(consts.Port)

	botKey := use_case.GetEnv(consts.BotKey)

	MyChatId, err := strconv.ParseInt(use_case.GetEnv(consts.ChatId), 10, 64)
	if err != nil {
		log.Println("CHAT_ID error:" + err.Error())
		return
	}

	//redisUrl := use_case.GetEnv(consts.RedisUrl)
	//scheduleRepository := repository.NewScheduleRepository(&redisUrl)
	//defer scheduleRepository.Close()

	bot, err := tgbotapi.NewBotAPI(botKey)
	if err != nil {
		log.Println(err)
		return
	}

	go httpserver.ServeHttpServer(&port)

	scheduleMap := make(map[string]*models.ScheduleEntity)
	if err != nil {
		log.Println(err)
		return
	}
	go delivery.CheckUpdates(bot, MyChatId, scheduleMap)
	delivery.RunPeriodically(bot, scheduleMap)

}
