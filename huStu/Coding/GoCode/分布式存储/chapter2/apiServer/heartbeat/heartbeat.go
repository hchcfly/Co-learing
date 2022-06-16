package heartbeat

import (
	"os"
	"strconv"
	"sync"
	"time"
	"lib/rabbitmq"
)
//  缓存所有的数据服务节点
var dataServers = make(map[string]time.Time)
var mutex sync.Mutex

func ListenHeartbeat() {
	q := rabbitMQ.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q,bind("apiServers")
	c := q.Consume()
	//  移除超时心跳包的数据服务器
	go removeExpiredDataServer()
	for msg := range c {
		dataServer ,e := strconv.Unquote(string(msg.Body))
		if e!=nil {
			panic(e)
		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

func removeExpiredDataServer() []string {
	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for s, t := range dataServers {
			//  清楚超过10s没有收到心跳消息的数据服务节点
			if t.Add(10*time.Second).Before(time.Now()) {
				
			}
		}
	}

}