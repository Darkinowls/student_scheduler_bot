package tests

import (
	"log"
	models "studentBot/features/telegram_bot/models"
	v "studentBot/features/telegram_bot/use_case"
	"testing"
)

func TestDTO(t *testing.T) {

	dto := models.ClientScheduleDTO{
		ScheduleFirstWeek: []models.ClientScheduleWeek{
			{
				Day: "Пн",
				Pairs: []models.ClientPair{
					{
						Time:  "10.00",
						Name:  "John Doe",
						Place: nil,
					},
					{
						Time:  "10.00",
						Name:  "Math",
						Place: nil,
					},
					{
						Time:  "12.00",
						Name:  "Jane Smith",
						Place: nil,
					},
				},
			},
			//{
			//	Day: "Вв",
			//	Pairs: []models.ClientPair{
			//		{
			//			Time:  "9.00",
			//			Name:  "John Doe",
			//			Place: &place,
			//		},
			//		{
			//			Time:  "16.00",
			//			Name:  "Jane Smith",
			//			Place: &place,
			//		},
			//	},
			//},
			// Add more weeks and pairs here.
		},
	}

	const chatId int64 = 123
	sMap, err := v.ParseClientScheduleDTOToScheduleEntities(&dto, chatId)
	if err != nil {
		log.Println("Validation failed:", err)
		t.Fail()
	} else {
		log.Println("Validation successful!")
		log.Println(*sMap["1:1:10.00"])
	}

}
