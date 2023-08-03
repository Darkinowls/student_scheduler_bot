package use_case

import "regexp"

var DayMap = map[string]int{
	"Пн": 1,
	"Вв": 2,
	"Ср": 3,
	"Чт": 4,
	"Пт": 5,
	"Сб": 6,
}

const MimeJson = "application/json"

var TimeRegex, _ = regexp.Compile(`^([01]?[0-9]|2[0-3])\.\d{2}$`)
