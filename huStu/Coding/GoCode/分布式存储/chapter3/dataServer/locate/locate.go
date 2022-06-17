package locate

import(
	"os"
	"fmt"
	"../../../lib/rabbitmq"
	"strconv"
)

//  判断文件是否在本地磁盘
func Locate(name string) bool {
	//  访问磁盘上对应的文件名
	_,err := os.Stat(name)
	//  判断文件名是否存在
	return !os.IsNotExist(err)
}

//  定位服务反馈
func StartLocate() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("dataServers")
	//  返回一个channel,用来接收消息
	c := q.Consume()
	for msg := range c {
		fmt.Println("Body :",msg.Body)
		object,e := strconv.Unquote(string(msg.Body))
		fmt.Println("Unquote Body :",msg.Body)
		if e != nil {
			panic(e)
		}
		//  对象存在本地磁盘上
		if Locate(os.Getenv("STORAGE_ROOT")+"/objects/" + object) {
			//  将HTTP监听接口发给临时消息队列
			q.Send(msg.ReplyTo,os.Getenv("LISTEN_ADDRESS"))
		}
	}


}