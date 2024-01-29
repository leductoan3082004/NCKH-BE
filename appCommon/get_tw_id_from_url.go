package appCommon

import (
	"fmt"
	"net/url"
	"strings"
)

func GetTweetID(source string) (string, error) {
	parsedURL, err := url.Parse(source)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %v", err)
	}

	parts := strings.Split(parsedURL.Path, "/")
	if len(parts) < 4 {
		return "", fmt.Errorf("invalid Twitter URL: no Tweet ID found")
	}

	return parts[3], nil
}
