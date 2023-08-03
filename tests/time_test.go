package tests

import (
	"log"
	"studentBot/features/telegram_bot/use_case"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	log.Println(use_case.GetKeyByTime(time.Now()))
}

func TestTimeNextWeek(t *testing.T) {
	log.Println(use_case.GetKeyByTime(time.Now().AddDate(0, 0, 7)))
}

func TestSleep(t *testing.T) {
	use_case.SleepIfNeeded(time.Now())
}
