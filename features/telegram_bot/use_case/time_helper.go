package use_case

import (
	"fmt"
	"time"
)

func GetKeyByTime(cTime time.Time) string {
	// Get the week number (odd/even)
	_, week := cTime.ISOWeek()
	weekNumber := week%2 == 0 // TODO : Impelement
	// Get the day of the week (number)
	dayNumber := int(cTime.Weekday())
	// Get the time (hours and minutes)
	timeString := cTime.Format("15.04")
	// Combine the components and format the final string
	return fmt.Sprintf("%d:%d:%s", weekNumber, dayNumber, timeString)
}

//func Sleep(cTime time.Time, duration time.Duration) {
//
//	time.Sleep()
//}
