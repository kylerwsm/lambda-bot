package util

import (
	"encoding/json"

	"github.com/kylerwsm/lambda-bot/pkg/entity"
)

var update *entity.Update

// SetUpdateFromBody creates and set an update struct from the request body.
func SetUpdateFromBody(body string) error {
	err := json.Unmarshal([]byte(body), &update)
	return err
}

// GetChatID extracts the user's chat ID from the update struct.
func GetChatID() int {
	return update.Message.Chat.ID
}

// GetText extracts the text message from the update struct.
func GetText() string {
	return update.Message.Text
}
