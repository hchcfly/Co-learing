package heartbeat
import (
	"os"
	"time"
	"lib/rabbitmq"
)
func StartHeartbeat() {
	//  创建一个rabbitMQ结构体
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	for {
		//  向apiServers发送本节点监听地址(心跳数据包)
		q.Publish("apiServers",os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5*time.Second)
	}

}