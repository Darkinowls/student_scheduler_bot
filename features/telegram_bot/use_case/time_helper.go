package use_case

import (
	"fmt"
	"log"
	"time"
)

func GetKeyByTime(cTime time.Time) string {
	// Get the week number (odd/even)
	_, week := cTime.ISOWeek()
	weekNumber := 1
	if week%2 == 0 {
		weekNumber = 2
	}

	// Get the day of the week (number)
	dayNumber := int(cTime.Weekday())
	// Get the time (hours and minutes)
	timeString := cTime.Format("15.04")
	// Combine the components and format the final string
	return fmt.Sprintf("%d:%d:%s", weekNumber, dayNumber, timeString)
}

func SleepIfNeeded(currentTime time.Time) {
	sleepDayIfDayOfWeek(currentTime, SunValue)
	sleepIfTimeOfDay(currentTime, HourToSleepDown, HourToWakeUp)
}

func sleepIfTimeOfDay(currentTime time.Time, hourToSleep int, hourToWakeUp int) {
	// Calculate the next sleep time on the same day as the current time
	nextSleepTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), hourToSleep, 0, 0, 0, currentTime.Location())
	// Calculate the next wake-up time on the same day as the current time
	nextWakeUpTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), hourToWakeUp, 0, 0, 0, currentTime.Location())

	if hourToWakeUp-hourToSleep <= 0 { // if wake-up time is tomorrow
		nextWakeUpTime = nextWakeUpTime.Add(DayInterval)
	}

	if currentTime.After(nextSleepTime) && currentTime.Before(nextWakeUpTime) {
		// Calculate the time duration between sleep and wake-up time
		duration := nextWakeUpTime.Sub(nextSleepTime)
		log.Println("Sleep Time:", nextSleepTime)
		log.Println("Wake Up Time:", nextWakeUpTime)
		log.Println("Duration:", duration)
		// Sleep until the wake-up time
		time.Sleep(duration)
	}
}

func sleepDayIfDayOfWeek(currentTime time.Time, targetDays ...time.Weekday) {
	weekDay := int(currentTime.Weekday())
	for day := range targetDays {
		if day == weekDay {
			time.Sleep(DayInterval)
		}
	}
}
