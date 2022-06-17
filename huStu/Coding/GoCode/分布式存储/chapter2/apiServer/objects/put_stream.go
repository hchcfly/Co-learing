package objects

import (
	"../heartbeat"
	"fmt"
	"../../../lib/objectstream"
)


func putSreanm(object string) (*objectstream.PutStream,error) {
	//  获取一个随机的数据服务节点的地址
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil,fmt.Errorf("connot find ant dataServer")
	}
	return objectstream.NewPutStream(server,object),nil

}

