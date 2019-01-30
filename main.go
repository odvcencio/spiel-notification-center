package main

import "spiel/notification-center/server"

func main() {
	server.CreateAndListen()
	//go server.WaitForMessages()
}
