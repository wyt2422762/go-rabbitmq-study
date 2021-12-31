package main

import (
	_ "go-rabbitmq-study/client"
	"go-rabbitmq-study/server"
	_ "go-rabbitmq-study/server"
)

func main() {
	//server.Send01()

	// client.Receive03()

	//client.Receive01()

	server.Send001()
}
