package services

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/kylerwsm/lambda-bot/pkg/util"
)

// SendResponse replies the user with the reply string.
func SendResponse(chatID int, reply string) error {
	env := "TELEGRAM_BOT_TOKEN"
	if token, ok := os.LookupEnv(env); ok {
		_, err := http.PostForm(
			fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token),
			url.Values{
				"chat_id":                  {strconv.Itoa(chatID)},
				"text":                     {reply},
				"disable_web_page_preview": {strconv.FormatBool(true)},
				"parse_mode":               {"Markdown"},
			})
		return err
	}
	return util.CreateEnvError(env)
}
