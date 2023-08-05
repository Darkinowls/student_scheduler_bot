package use_case

import (
	"errors"
	"fmt"
	"strings"
	"studentBot/features/telegram_bot/consts"
	"studentBot/features/telegram_bot/models"
)

func ParseClientScheduleDTOToScheduleEntities(dto *models.ClientScheduleDTO, chatId int64) (scheduleMap map[string]*models.ScheduleEntity, err error) {
	scheduleMap = make(map[string]*models.ScheduleEntity)
	for _, day := range dto.ScheduleFirstWeek {
		err = mapClientPairsToPairEntities(&day, chatId, scheduleMap, 1)
		if err != nil {
			return scheduleMap, err
		}
	}
	for _, day := range dto.ScheduleSecondWeek {
		err = mapClientPairsToPairEntities(&day, chatId, scheduleMap, 2)
		if err != nil {
			return scheduleMap, err
		}
	}
	return scheduleMap, nil
}

// Function to convert ClientPair slice to PairEntity slice
func mapClientPairsToPairEntities(day *models.ClientScheduleWeek, chatId int64, scheduleMap map[string]*models.ScheduleEntity, weekNum int) error {

	if _, f := consts.DayMap[day.Day]; !f {
		return errors.New("Day field is incorect")
	}

	for _, pair := range day.Pairs {

		if err := validateClientPair(pair); err != nil {
			return err
		}

		pairEntity := models.PairEntity{
			Name:   pair.Name,
			ChatID: chatId,
		}

		if pair.Place != nil {
			pairEntity.Link = *pair.Place
		}

		key := makeDatetime(weekNum, consts.DayMap[day.Day], pair.Time)
		if value, found := scheduleMap[key]; found {
			pairs := &value.Pairs
			value.Pairs = append(*pairs, pairEntity)
			continue
		}
		value := models.ScheduleEntity{Pairs: []models.PairEntity{pairEntity}}
		scheduleMap[key] = &value
	}
	return nil
}

// Validator function to validate the ClientPair
func validateClientPair(pair models.ClientPair) error {
	if !consts.TimeRegex.MatchString(pair.Time) {
		return errors.New("Time field is incorect")
	}

	// Add additional checks for time format, if needed.

	if pair.Name == "" {
		return errors.New("Name field is required")
	}

	if pair.Place != nil && strings.TrimSpace(*pair.Place) == "" {
		return errors.New("Place cannot be an empty string")
	}

	return nil
}

func makeDatetime(week int, day int, time string) string {
	return fmt.Sprintf("%d:%d:%s", week, day, time)
}
