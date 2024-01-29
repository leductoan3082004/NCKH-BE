package appCommon

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type ImageInfo struct {
	ImageData []byte
	Ext       string
}

func GetImageFromUrl(url string) (*ImageInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, ErrInvalidRequest(err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrInvalidRequest(errors.New("the link is error"))
	}
	defer resp.Body.Close()
	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		return nil, ErrInvalidRequest(errors.New("invalid image url"))
	}

	extension := contentType[6:]
	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrInternal(err)
	}
	return &ImageInfo{
		ImageData: imageData,
		Ext:       extension,
	}, nil
}
