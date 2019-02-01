package main

import (
	"spiel/notification-center/messaging"
	"spiel/notification-center/server"
)

func main() {
	messaging.SubscribeToTopics()
	server.CreateAndListen()
}
