package appCommon

import "net/url"

func GetPathFromUrl(urlString string) (string, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "", ErrInvalidRequest(err)
	}
	res := S3Mindzone + parsedURL.Path

	return res, nil
}
