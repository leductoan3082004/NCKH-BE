package appCommon

import "encoding/json"

func ToJson(payload interface{}) ([]byte, error) {
	byteData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return byteData, nil
}
