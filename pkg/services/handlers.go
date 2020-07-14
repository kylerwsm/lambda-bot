package services

import (
	"github.com/kylerwsm/lambda-bot/pkg/repositories"
	"github.com/kylerwsm/lambda-bot/pkg/util"
)

// PingHandler handles ping commands.
func PingHandler() string {
	return "pong"
}

// ShortLinksHandler handles shortlinks commands.
func ShortLinksHandler() (string, error) {
	links, err := repositories.GetShortLinks()
	if err != nil {
		return "", nil
	}
	linkString := "These are your shortened links.\n\n" + util.Join(links)
	return linkString, nil
}

// DefaultHandler handles unknown commands.
func DefaultHandler() string {
	return (`
	I am Kyler Bot and I can help you manage chores.

	Please use the following commands:
	/ping - replies pong
	/shortlinks - lists your shortened links
	`)
}
