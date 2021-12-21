package config

import (
	"fmt"

	"github.com/streadway/amqp"
)


var Conn *amqp.Connection
var err error

func init(){
	//连接rabbitmq
	Conn, err = amqp.Dial("amqp://test:test@47.105.71.238:5672/my_vhost")
	if err != nil {
		fmt.Println("连接消息队列失败")
		panic("连接消息队列失败")
	}
}