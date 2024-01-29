package appCommon

import "encoding/json"

func StructToMap(data interface{}) (map[string]interface{}, error) {
	res, err := json.Marshal(data)
	if err != nil {
		return nil, ErrInternal(err)
	}
	mapData := make(map[string]interface{})
	if err := json.Unmarshal([]byte(res), &mapData); err != nil {
		return nil, ErrInternal(err)
	}
	return mapData, nil
}
