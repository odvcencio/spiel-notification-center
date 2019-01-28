package server

import (
	"spiel/notification-center/messaging"
	"sync"

	"github.com/nsqio/go-nsq"

	"github.com/labstack/echo"
)

// CreateAndListen creates main echo server and
// starts listening on specific port
func CreateAndListen() {
	e := echo.New()

	// Connect logger
	e.Use(connectLogger())

	// Registering routes
	registerRoutes(e)

	e.Start(":8080")
}

// WaitForMessages waits for messages sent through nsq
func WaitForMessages() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	// Topic: Question To User
	messaging.GetConsumer(
		messaging.TopicQuestionToUser,
		nsq.HandlerFunc(handleTopicQuestionToUser),
	)

	// Blocking main thread
	wg.Wait()
}
