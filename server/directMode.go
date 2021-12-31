package server

import (
	"bufio"
	"fmt"
	"github.com/streadway/amqp"
	"go-rabbitmq-study/config"
	"os"
)

//rabbitmq - 直连模式 - 发送端
func Send01() {
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
		err = ch.Publish("", DEMO_QUEUE_NAME, false, false, amqp.Publishing{
			Body:        byt,
			ContentType: "text/plain",
		})
		if err != nil {
			fmt.Println("消息发送失败")
		}

	}
}

func Send001() {
	//连接rabbitmq
	conn := config.Conn
	defer conn.Close()

	//打开通道
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("打开通道失败")
	}
	defer ch.Close()

	//开启confirm
	err = ch.Confirm(false)
	if err != nil {
		fmt.Println("开启confirm失败")
	}
	confirm := ch.NotifyPublish(make(chan amqp.Confirmation, 1))

	//开启return
	ret := ch.NotifyReturn(make(chan amqp.Return, 1))

	reader := bufio.NewReader(os.Stdin) //读取输入的内容

	for {
		byt, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("读取输入失败")
		}

		//发送消息
		err = ch.Publish("amq.direct", DEMO_QUEUE_NAME, true, false, amqp.Publishing{
			Body:        byt,
			ContentType: "text/plain",
		})
		if err != nil {
			fmt.Println("消息发送失败")
		}

		go listenConfirm(confirm) // confirm处理方法

		go listenReturn(ret) // return处理方法
	}
}

// 消息confirm方法
func listenConfirm(confirms <-chan amqp.Confirmation) {
	if confirmed := <-confirms; confirmed.Ack {
		fmt.Println("消息发送成功")
	} else {
		fmt.Println("消息发送失败")
	}
}

//消息return方法
func listenReturn(ret <-chan amqp.Return) {
	fmt.Println("消息发送失败，找不到对应的队列，返回消息")
	re := <-ret
	fmt.Println("replyCode = ", re.ReplyCode)
	fmt.Println("replyText = ", re.ReplyText)
	fmt.Println("exchange = ", re.Exchange)
	fmt.Println("routingKey = ", re.RoutingKey)
	fmt.Println("msg = ", string(re.Body))
}
