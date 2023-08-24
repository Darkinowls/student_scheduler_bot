package delivery

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"studentBot/features/telegram_bot/consts"
	"studentBot/features/telegram_bot/repository"
	"studentBot/features/telegram_bot/use_case"
	"time"
)

func RunPeriodically(bot *tgbotapi.BotAPI, scheduleRepo repository.ScheduleRepository) {
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
			sendScheduleInTime(&currentTime, scheduleRepo, bot)
		default:
			currentTime := time.Now()
			sendScheduleInTime(&currentTime, scheduleRepo, bot)
		}

	}
}

func CheckUpdates(bot *tgbotapi.BotAPI, chatId int64, scheduleRepository repository.ScheduleRepository) {

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
		err = scheduleRepository.DeleteAllRecords()
		if err != nil {
			log.Println(err)
			continue
		}
		err = scheduleRepository.SaveScheduleEntities(sMap)
		if err != nil {
			log.Println(err)
			continue
		}
		suchMessage := "The schedule is successfully set "
		log.Print(suchMessage, sMap)
		msg := tgbotapi.NewMessage(chatId, suchMessage)
		bot.Send(msg)
	}
}

func sendScheduleInTime(currentTime *time.Time, scheduleRepo repository.ScheduleRepository, bot *tgbotapi.BotAPI) {
	keys := use_case.GetKeysByTime(currentTime, consts.RunScheduleMinute)
	scheduleMap, err := scheduleRepo.GetScheduleEntities()
	if err != nil {
		log.Println(err)
		return
	}
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
