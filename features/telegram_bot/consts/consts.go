package consts

import (
	"time"
)

const (
	MimeJson          = "application/json"
	HourToSleepDown   = 18
	HourToWakeUp      = 8
	RunScheduleMinute = 2
	DayInterval       = 24 * time.Hour
)

var (
	// DefaultTimezone is the default time zone for the package.
	DefaultTimezone *time.Location
)
