package tests

import (
	delivery "studentBot/features/telegram_bot/delivery"
	models "studentBot/features/telegram_bot/models"
	"testing"
)

func TestRequests(t *testing.T) {
	dto, err := models.UnmarshalServerScheduleDto(delivery.GetRequests())
	if err != nil {
		panic("TestRequests")
	}
	println(dto.Data.GroupCode)

}
