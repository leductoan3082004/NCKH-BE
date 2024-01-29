package appCommon

import (
	"context"
	"mindzone/plugin/pubsub"
)

type HandlerFunc func(ctx context.Context, message *pubsub.Message, workerId int) error
type SubJob struct {
	Title           string
	NumWorker       int
	Topic           string
	TopicDeadLetter string
	Hdl             HandlerFunc
}

func NewSubJob(title, topic, topicDeadLetter string, numWorker int, hdl HandlerFunc) *SubJob {
	return &SubJob{
		Title:           title,
		NumWorker:       numWorker,
		Topic:           topic,
		TopicDeadLetter: topicDeadLetter,
		Hdl:             hdl,
	}
}
