package repository

import (
	"studentBot/features/telegram_bot/models"
)

type InMemoryScheduleRepository struct {
	scheduleMap map[string]*models.ScheduleEntity
}

func (r *InMemoryScheduleRepository) DeleteAllRecords() error {
	r.scheduleMap = make(map[string]*models.ScheduleEntity)
	return nil
}

func (r *InMemoryScheduleRepository) SaveScheduleEntities(sMap map[string]*models.ScheduleEntity) error {
	for key, value := range sMap {
		r.scheduleMap[key] = value
	}
	return nil
}

func (r *InMemoryScheduleRepository) GetScheduleEntities() (map[string]*models.ScheduleEntity, error) {
	return r.scheduleMap, nil
}

func (r *InMemoryScheduleRepository) Close() error {
	// In this in-memory version, there is no need to close any connection.
	return nil
}

func NewInMemoryScheduleRepository() ScheduleRepository {
	return &InMemoryScheduleRepository{
		scheduleMap: make(map[string]*models.ScheduleEntity),
	}
}
