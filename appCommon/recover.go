package appCommon

import "fmt"

func Recover() {
	if r := recover(); r != nil {
		fmt.Println("Recovered error:", r)
	}
}
