package services

import (
	"fmt"
	"strings"

	"github.com/kylerwsm/lambda-bot/pkg/entity"
	"github.com/kylerwsm/lambda-bot/pkg/repositories"
)

// PingHandler handles ping commands.
func PingHandler() string {
	return "pong"
}

// Join joins and stringifies the link slice.
func Join(links []entity.Link) string {
	formattedLinks := make([]string, 0)
	for _, link := range links {
		formattedLink := fmt.Sprintf("*%s*\n→ %s", link.ShortLink, link.OriginalLink)
		formattedLinks = append(formattedLinks, formattedLink)
	}
	return strings.Join(formattedLinks, "\n\n")
}

// ShortLinksHandler handles shortlinks commands.
func ShortLinksHandler() (string, error) {
	links, err := repositories.GetShortLinks()
	if err != nil {
		return "", nil
	}
	linkString := "These are your shortened links.\n\n" + Join(links)
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
