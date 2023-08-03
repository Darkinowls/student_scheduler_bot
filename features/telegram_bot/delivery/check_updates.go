package delivery

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"studentBot/features/telegram_bot/models"
	"studentBot/features/telegram_bot/use_case"
)

func CheckUpdates(bot *tgbotapi.BotAPI, chatId int64, scheduleMap map[string]*models.ScheduleEntity) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil || update.FromChat().ID != chatId {
			continue
		} // If I got a message

		dto, err := use_case.GetClientScheduleDTOFromUpdate(&update, bot)
		if err != nil {
			msg := tgbotapi.NewMessage(chatId, err.Error())
			bot.Send(msg)
			log.Println(err)
			continue
		}

		sMap, err := use_case.ParseClientScheduleDTOToScheduleEntities(dto, chatId)
		if err != nil {
			msg := tgbotapi.NewMessage(chatId, err.Error())
			bot.Send(msg)
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
