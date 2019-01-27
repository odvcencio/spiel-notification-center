package server

import (
	"fmt"
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

	messaging.GetConsumer(
		messaging.TopicQuestionFromUser,
		nsq.HandlerFunc(func(message *nsq.Message) error {
			fmt.Println(string(message.Body))
			return nil
		}),
	)

	// Blocking main thread
	wg.Wait()
}
