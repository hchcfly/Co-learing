package locate

import (
	"../../../lib/rabbitmq"
	"os"
	"strconv"
	"time"
)


func Locate(name string) string {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	q.Publish("dataServers",name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	//  阻塞等待数据服务节点向自己发送的反馈消息
	msg := <- c
	s , _ := strconv.Unquote(string(msg.Body))
	return s
}

func Exist(name string) bool {
	return Locate(name) != ""
}

