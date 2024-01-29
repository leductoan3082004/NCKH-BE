package appCommon

import (
	"encoding/json"
)

type chatCompletion struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func TakeSummaryResult(data interface{}) (*chatCompletion, error) {
	res, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var ans chatCompletion
	if err := json.Unmarshal(res, &ans); err != nil {
		return nil, err
	}
	return &ans, nil
}
