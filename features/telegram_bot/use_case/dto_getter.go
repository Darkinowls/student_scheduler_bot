package use_case

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"studentBot/features/telegram_bot/models"
)

func GetClientScheduleDTOFromUpdate(update *tgbotapi.Update, bot *tgbotapi.BotAPI) (dto *models.ClientScheduleDTO, err error) {
	downloadedFile, err := downloadIncomingFile(update, bot)
	defer downloadedFile.Close()
	if err != nil {
		return nil, errors.New("Cannot load the file")
	}
	dto, err = getClientScheduleDTOByFile(downloadedFile)
	if err != nil {
		return nil, errors.New("Json syntaxis error")
	}
	return dto, nil
}

func getClientScheduleDTOByFile(downloadedFile *os.File) (*models.ClientScheduleDTO, error) {
	_, err := downloadedFile.Stat()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(downloadedFile.Name())
	return models.UnmarshalClientScheduleDTO(data)
}

func downloadIncomingFile(update *tgbotapi.Update, bot *tgbotapi.BotAPI) (downloadedFile *os.File, err error) {
	if update.Message.Document.MimeType != MimeJson {
		log.Println("File is not a json file")
		return
	}

	// Download the file content
	resp, err := bot.GetFileDirectURL(update.Message.Document.FileID)
	if err != nil {
		log.Println("Error getting file URL:", err)
		return
	}

	// Create a new file to save the downloaded content
	downloadedFile, err = os.Create("cache.json")
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	//defer downloadedFile.Close()

	// Download the file content from the obtained URL
	fileContent, err := http.Get(resp)
	if err != nil {
		log.Println("Error downloading file content:", err)
		return
	}
	defer fileContent.Body.Close()

	// Save the file content to the new file
	_, err = io.Copy(downloadedFile, fileContent.Body)
	if err != nil {
		log.Println("Error saving file content:", err)
		return
	}

	log.Println("File downloaded successfully!")
	return downloadedFile, nil
}
