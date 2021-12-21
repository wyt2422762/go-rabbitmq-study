package server

import (
	"bufio"
	"fmt"
	"github.com/streadway/amqp"
	"os"
	"go-rabbitmq-study/config"
)


//rabbitmq - topic模式 - 发送端
func Send03() {
	//连接rabbitmq
	conn := config.Conn
	defer conn.Close()

	//打开通道
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("打开通道失败")
	}
	defer ch.Close()

	reader := bufio.NewReader(os.Stdin) //读取输入的内容

	for {
		byt, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("读取输入失败")
		}

		//发送消息
		err = ch.Publish(TOPIC_EXCHANGE_NAME, TOPIC1, false, false, amqp.Publishing{
			Body:        byt,
			ContentType: "text/plain",
		})
		if err != nil {
			fmt.Println("消息发送失败")
		}

	}
}
