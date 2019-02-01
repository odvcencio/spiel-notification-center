package messaging

import (
	nsq "github.com/nsqio/go-nsq"
)

// SubscribeToTopics waits for messages sent through nsq
func SubscribeToTopics() {
	// Topic: Question To User
	getConsumer(
		topicQuestionToUser,
		nsq.HandlerFunc(handleTopicQuestionToUser),
	)
}
