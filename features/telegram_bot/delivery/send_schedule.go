package delivery

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"studentBot/features/telegram_bot/models"
	"time"
)

func RunPeriodically(interval time.Duration, bot *tgbotapi.BotAPI, scheduleMap map[string]*models.ScheduleEntity, chatId int64) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		msg := tgbotapi.NewMessage(chatId, "Periodical")
		bot.Send(msg)
	}
}
