package delivery

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"studentBot/features/telegram_bot/consts"
	"studentBot/features/telegram_bot/models"
	"studentBot/features/telegram_bot/use_case"
	"time"
)

func RunPeriodically(bot *tgbotapi.BotAPI, scheduleMap map[string]*models.ScheduleEntity) {
	minuteTicker := time.NewTicker(consts.RunScheduleMinute * time.Minute)
	defer minuteTicker.Stop()
	hourTicker := time.NewTicker(time.Hour)
	defer hourTicker.Stop()
	for {
		<-minuteTicker.C
		select {
		case <-hourTicker.C:
			currentTime := time.Now()
			use_case.SleepIfNeeded(currentTime)
			sendScheduleInTime(&currentTime, scheduleMap, bot)
		default:
			currentTime := time.Now()
			sendScheduleInTime(&currentTime, scheduleMap, bot)
		}

	}
}

func CheckUpdates(bot *tgbotapi.BotAPI, chatId int64, scheduleMap map[string]*models.ScheduleEntity) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil || update.FromChat().ID != chatId {
			continue
		} // If I got a message in the chat

		dto, err := use_case.GetClientScheduleDTOFromUpdate(&update, bot)
		if err != nil {
			msg := tgbotapi.NewMessage(chatId, err.Error())
			_, err2 := bot.Send(msg)
			if err2 != nil {
				log.Println(err2)
				continue
			}
			log.Println(err)
			continue
		}

		sMap, err := use_case.ParseClientScheduleDTOToScheduleEntities(dto, chatId)
		if err != nil {
			msg := tgbotapi.NewMessage(chatId, err.Error())
			_, err2 := bot.Send(msg)
			if err2 != nil {
				log.Println(err2)
				continue
			}
			log.Println(err)
			continue
		}
		for k := range scheduleMap {
			delete(scheduleMap, k)
		}
		for k, v := range sMap {
			scheduleMap[k] = v
		}
		log.Print(sMap)

	}
}

func sendScheduleInTime(currentTime *time.Time, scheduleMap map[string]*models.ScheduleEntity, bot *tgbotapi.BotAPI) {
	keys := use_case.GetKeysByTime(currentTime, consts.RunScheduleMinute)
	for _, key := range keys {
		schedule, found := scheduleMap[key]
		if !found {
			continue
		}
		for _, p := range schedule.Pairs {
			msg := tgbotapi.NewMessage(p.ChatID, p.Name+"\n"+p.Link)
			_, err := bot.Send(msg)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
}
