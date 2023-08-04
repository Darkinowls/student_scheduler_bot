package delivery

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"studentBot/features/telegram_bot/models"
	"studentBot/features/telegram_bot/use_case"
	"time"
)

func RunPeriodically(bot *tgbotapi.BotAPI, scheduleMap map[string]*models.ScheduleEntity) {
	ticker := time.NewTicker(use_case.RunScheduleMinute * time.Minute)
	defer ticker.Stop()

	// TODO: check if sleep every hour and not 2 minutes

	for {
		<-ticker.C
		currentTime := time.Now()
		use_case.SleepIfNeeded(currentTime)
		keys := use_case.GetKeysByTime(currentTime, use_case.RunScheduleMinute)
		for _, key := range keys {
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
}
