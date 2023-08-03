package tests

import (
	"log"
	"studentBot/features/telegram_bot/use_case"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	log.Println(use_case.GetKeysByTime(time.Now(), 1))
}

func TestTimeNextWeek(t *testing.T) {
	log.Println(use_case.GetKeysByTime(time.Now().AddDate(0, 0, 7), 1))
}

func TestSleep(t *testing.T) {
	use_case.SleepIfNeeded(time.Now())
}
