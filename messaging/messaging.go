package messaging

import (
	"log"
	"os"

	"github.com/nsqio/go-nsq"
)

// myChannelName indicates channel name for this
// service
var myChannelName = "notification-center"

// A map of topic => consumer
var consumers = map[string]*nsq.Consumer{}

// GetConsumer first checks for any existing consumer
// for the specified topic, if doesn't exist, creates
// one and returns
func GetConsumer(topic string, handlers ...nsq.Handler) (*nsq.Consumer, error) {
	if q, ok := consumers[topic]; ok {
		return q, nil
	}

	// Create new consumer
	config := nsq.NewConfig()
	q, err := nsq.NewConsumer(topic, myChannelName, config)
	if err != nil {
		return nil, err
	}

	// Add Handler
	for _, handler := range handlers {
		q.AddHandler(handler)
	}

	// Connecting to nsqlookupd
	if err := q.ConnectToNSQD(os.Getenv("NSQD_HOST") + ":" + os.Getenv("NSQD_PORT")); err != nil {
		return nil, err
	}
	log.Println("connected to nsqd successfully")

	// Saving consumer
	consumers[topic] = q

	return q, nil
}
