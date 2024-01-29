package appCommon

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

func HMACEncode(data interface{}, secret string) (string, error) {
	byteData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(byteData)

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
