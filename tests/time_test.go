package tests

import (
	"log"
	"studentBot/features/telegram_bot/consts"
	"studentBot/features/telegram_bot/use_case"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	tim := time.Now().In(consts.DefaultTimezone)
	log.Println(use_case.GetKeysByTime(&tim, 1))
}

func TestTimeNextWeek(t *testing.T) {
	tim := time.Now().In(consts.DefaultTimezone).AddDate(0, 0, 7)
	log.Println(use_case.GetKeysByTime(&tim, 1))
}

func TestSleep(t *testing.T) {
	use_case.SleepIfNeeded(time.Now().In(consts.DefaultTimezone))
}

func TestGetKeyByTime(t *testing.T) {
	data := use_case.GetKeyByTime(time.Now().In(consts.DefaultTimezone))
	log.Println(data)
}
