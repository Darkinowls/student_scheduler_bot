package delivery

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"studentBot/features/telegram_bot/models"
	"studentBot/features/telegram_bot/use_case"
	"time"
)

func RunPeriodically(interval time.Duration, bot *tgbotapi.BotAPI, scheduleMap map[string]*models.ScheduleEntity, chatId int64) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		key := use_case.GetKeyByTime(time.Now())
		schedule, found := scheduleMap[key]
		if !found {
			continue
		}
		for _, p := range schedule.Pairs {
			msg := tgbotapi.NewMessage(p.ChatID, p.Name+"\n"+p.Link)
			bot.Send(msg)
		}
	}
}
