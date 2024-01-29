package appCommon

import (
	"encoding/json"
	"strings"
)

type Node struct {
	Children []*Node `json:"children"`
	Content  []struct {
		Text string `json:"text"`
	} `json:"content"`
}

func dfsLinkedList(root *Node, arr *[]string) {
	if root == nil {
		return
	}
	for i := range root.Content {
		*arr = append(*arr, root.Content[i].Text)
	}
	for i := range root.Children {
		dfsLinkedList(root.Children[i], arr)
	}
}
func GetText1(s string) (string, error) {
	var res []*Node
	if err := json.Unmarshal([]byte(s), &res); err != nil {
		return "", err
	}
	var ans []string

	for i := range res {
		dfsLinkedList(res[i], &ans)
	}
	return strings.Join(ans, " "), nil
}
