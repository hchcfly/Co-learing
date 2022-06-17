package heartbeat

import (
	"os"
	"strconv"
	"sync"
	"time"
	"math/rand"
	"../../../lib/rabbitmq"
)
//  缓存所有的数据服务节点
var dataServers = make(map[string]time.Time)
var mutex sync.Mutex


//  监听数据节点发送的心跳包
func ListenHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("apiServers")
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
//  移除过期的数据节点
func removeExpiredDataServer() []string {
	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for s, t := range dataServers {
			//  清楚超过10s没有收到心跳消息的数据服务节点
			if t.Add(10*time.Second).Before(time.Now()) { //add 10 是否在Now之前
				delete(dataServers,s)
			}
		}
		mutex.Unlock()
	}
}
//  获取当前所有数据节点
func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string,0)
	for s, _ := range dataServers {
		ds = append(ds,s)
	}
	return ds
}
//  在当前所有的数据节点随机选取一个节点返回
func ChooseRandomDataServer() string {
	ds := GetDataServers()
	n := len(ds)
	if n == 0{
		return ""
	}
	return ds[rand.Intn(n)]
}
