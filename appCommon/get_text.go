package appCommon

import (
	"encoding/json"
	"strings"
)

func dfs(k interface{}, arr *[]string) {
	value, ok := k.(map[string]interface{})
	if ok {
		for key := range value {
			//fmt.Println(key, reflect.TypeOf(value[key]))

			if key == "text" {
				*arr = append(*arr, value[key].(string))
			}
			dfs(value[key], arr)
		}
	}
	val, ok := k.([]interface{})

	if ok {
		for i := range val {
			value, ok := val[i].(map[string]interface{})
			if ok {
				for key := range value {
					//fmt.Println(key, reflect.TypeOf(value[key]))

					if key == "text" {
						*arr = append(*arr, value[key].(string))
					}
					dfs(value[key], arr)
				}
			}
		}
	}
}
func solve(k interface{}) []string {
	var res []string
	dfs(k, &res)
	return res
}

func GetText(s string) string {
	var tmp interface{}
	if err := json.Unmarshal([]byte(s), &tmp); err != nil {
		panic(err)
	}
	arr := solve(tmp)
	return strings.Join(arr, " ")
}
