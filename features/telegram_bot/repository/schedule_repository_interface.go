package repository

import (
	"studentBot/features/telegram_bot/models"
)

type ScheduleRepository interface {
	DeleteAllRecords() error
	SaveScheduleEntities(sMap map[string]*models.ScheduleEntity) error
	GetScheduleEntities() (map[string]*models.ScheduleEntity, error)
	Close() error
}
