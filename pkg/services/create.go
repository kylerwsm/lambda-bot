package services

import "github.com/kylerwsm/lambda-bot/pkg/entity"

// CreateResponse creates a response to the user.
func CreateResponse(command string) (string, error) {
	var resp string
	var err error
	switch command {
	case entity.Ping:
		resp = PingHandler()
	case entity.ShortLinks:
		resp, err = ShortLinksHandler()
	default:
		resp = DefaultHandler()
	}
	return resp, err
}
