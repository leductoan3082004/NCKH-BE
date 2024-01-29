package appCommon

import (
	"fmt"
	"net/http"
	"strings"
)

func SendNotiTelegram(Title string, lineOfText ...string) {
	var message string
	message += Title + "\n---\n"
	for _, line := range lineOfText {
		message += line + "\n"
	}
	url := "https://api.telegram.org/bot6069825434:AAHJdfXND0IAgLWX6lX0Il8NG_IbpbP_9_A/sendMessage"
	method := "POST"

	payload := strings.NewReader(`{
	  "text": "` + message + `",
	  "chat_id": "-1002000854996",
	  "parse_mode": "Markdown"
  }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
}
