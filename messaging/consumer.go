package messaging

import (
	"log"
	"os"

	"github.com/nsqio/go-nsq"
)

// channelName indicates channel name for this
// service
var channelName = "notification-center"

// A map of topic => consumer
var consumers = map[string]*nsq.Consumer{}

// getConsumer first checks for any existing consumer
// for the specified topic, if doesn't exist, creates
// one and returns
func getConsumer(topic string, handlers ...nsq.Handler) (*nsq.Consumer, error) {
	if q, ok := consumers[topic]; ok {
		return q, nil
	}

	// Create new consumer
	config := nsq.NewConfig()
	q, err := nsq.NewConsumer(topic, channelName, config)
	if err != nil {
		return nil, err
	}

	// Add Handler
	for _, handler := range handlers {
		q.AddHandler(handler)
	}

	// Connecting to nsqlookupd
	if err := q.ConnectToNSQLookupd(os.Getenv("NSQLOOKUPD_HOST") + ":" + os.Getenv("NSQLOOKUPD_PORT")); err != nil {
		return nil, err
	}
	log.Println("connected to nsqlookupd successfully")

	// Saving consumer
	consumers[topic] = q

	return q, nil
}
