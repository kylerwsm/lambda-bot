package util

import (
	"fmt"
	"strings"

	"github.com/kylerwsm/lambda-bot/pkg/entity"
)

// Join joins and stringifies the link slice.
func Join(links []entity.Link) string {
	formattedLinks := make([]string, 0)
	for _, link := range links {
		formattedLink := fmt.Sprintf("• %s → %s", link.ShortLink, link.OriginalLink)
		formattedLinks = append(formattedLinks, formattedLink)
	}
	return strings.Join(formattedLinks, "\n")
}
