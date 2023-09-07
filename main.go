package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"strconv"
	"studentBot/features/telegram_bot/consts"
	"studentBot/features/telegram_bot/delivery"
	"studentBot/features/telegram_bot/repository"
	"studentBot/features/telegram_bot/use_case"
	"time"
)

func main() {

	_ = godotenv.Overload()

	botKey := use_case.GetEnv(consts.BotKey)

	redisUrl := use_case.GetEnv(consts.RedisUrl)

	loc, err := time.LoadLocation(use_case.GetEnv(consts.TZ))
	if err != nil {
		log.Println("TZ error:" + err.Error())
		return
	}
	consts.DefaultTimezone = loc

	MyChatId, err := strconv.ParseInt(use_case.GetEnv(consts.ChatId), 10, 64)
	if err != nil {
		log.Println("CHAT_ID error:" + err.Error())
		return
	}

	bot, err := tgbotapi.NewBotAPI(botKey)
	if err != nil {
		log.Println(err)
		return
	}

	scheduleRepo := repository.NewRedisScheduleRepository(&redisUrl)
	if err != nil {
		log.Println(err)
		return
	}
	go delivery.CheckUpdates(bot, MyChatId, scheduleRepo)
	delivery.RunPeriodically(bot, scheduleRepo)

}
