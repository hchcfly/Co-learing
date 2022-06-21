package locate

import(
	"os"
	"../../../lib/rabbitmq"
	"strconv"
	"sync"
)

var objects = make(map[string]int) // 缓存要定位的对象
var mutex sync.Mutex

// //  判断文件是否在本地磁盘
// func Locate(name string) bool {
// 	//  访问磁盘上对应的文件名
// 	_,err := os.Stat(name)
// 	//  判断文件名是否存在
// 	return !os.IsNotExist(err)
// }

func Locate(hash string) bool {
	mutex.Lock()
	_,ok := objects[hash]
	mutex.Unlock()
	return ok
}

func Add(hash string) {
	mutex.Lock()
	objects[hash] = 1;
	mutex.Unlock()
}

func Del(hash string) {
	mutex.Lock()
	delete(objects,hash)
	mutex.Unlock()
}

//  定位服务反馈 (直接读取缓存,不需要系统调用了)
func StartLocate() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("dataServers")
	//  返回一个channel,用来接收消息
	c := q.Consume()
	for msg := range c {
		//fmt.Println("Body :",msg.Body)
		hash,e := strconv.Unquote(string(msg.Body))
		//fmt.Println("Unquote Body :",msg.Body)
		if e != nil {
			panic(e)
		}
		exist := Locate(hash)
		//  对象存在本地磁盘上
		if exist {
			//  将HTTP监听接口发给临时消息队列
			q.Send(msg.ReplyTo,os.Getenv("LISTEN_ADDRESS"))
		}
	}
}

func CollectObjects() {
	files,_ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/*")
	for i, v := range files {
		hash := filepath.Base(files[i])
		objects[hash] = 1
	}
}