package services

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// SendResponse replies the user with the reply string.
func SendResponse(chatID int, reply string) error {
	if token, ok := os.LookupEnv("TELEGRAM_BOT_TOKEN"); ok {
		_, err := http.PostForm(
			fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token),
			url.Values{
				"chat_id":                  {strconv.Itoa(chatID)},
				"text":                     {reply},
				"disable_web_page_preview": {strconv.FormatBool(true)},
			})
		return err
	}
	return errors.New("TELEGRAM_BOT_TOKEN environment variable is not defined")
}
