package tests

import (
	"log"
	"studentBot/features/telegram_bot/use_case"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	tim := time.Now()
	log.Println(use_case.GetKeysByTime(&tim, 1))
}

func TestTimeNextWeek(t *testing.T) {
	tim := time.Now().AddDate(0, 0, 7)
	log.Println(use_case.GetKeysByTime(&tim, 1))
}

func TestSleep(t *testing.T) {
	use_case.SleepIfNeeded(time.Now())
}

func TestGetKeyByTime(t *testing.T) {
	data := use_case.GetKeyByTime(time.Now())
	log.Println(data)
}
