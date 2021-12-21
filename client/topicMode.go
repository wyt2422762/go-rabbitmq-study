package client

import (
	"fmt"
	"go-rabbitmq-study/config"
	"time"
)

//rabbitmq - topic模式 - 接收端
func Receive03() {
	//rabbitmq连接
	conn := config.Conn
	defer conn.Close()

	//打开通道
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("打开通道失败")
	}
	defer ch.Close()

	msgs, err := ch.Consume(TOPIC_QUEUE2, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println("读取数据失败")
	}

	for d := range msgs {
		fmt.Printf("收到消息: %s\n", string(d.Body))
		//手动应答(确认) multiple为true表示批量处理以前未应答的消息
		d.Ack(false)
		//手动应答(拒绝) requeue为true表示将消息重新放回队列，false表示不再重新入队，如果配置了死信队列则进入死信队列
		//d.Reject(false)
		//休眠5秒
		time.Sleep(time.Duration(5) * time.Second)
	}

	fmt.Println("程序退出")
}