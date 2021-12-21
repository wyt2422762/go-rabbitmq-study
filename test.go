package main

import (
	"go-rabbitmq-study/client"
	_ "go-rabbitmq-study/server"
)

func main() {
	//server.Send01()

	client.Receive03()
}
