package consts

const (
	SunKey   = "Нд"
	MonKey   = "Пн"
	TueKey   = "Вв"
	WedKey   = "Ср"
	ThuKey   = "Чт"
	FriKey   = "Пт"
	SatKey   = "Сб"
	SunValue = 0
	MonValue = 1
	TueValue = 2
	WedValue = 3
	ThuValue = 4
	FriValue = 5
	SatValue = 6
)

var DayMap = map[string]int{
	SunKey: SunValue,
	MonKey: MonValue,
	TueKey: TueValue,
	WedKey: WedValue,
	ThuKey: ThuValue,
	FriKey: FriValue,
	SatKey: SatValue,
}
