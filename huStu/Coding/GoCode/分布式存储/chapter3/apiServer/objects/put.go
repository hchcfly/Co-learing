package objects

import(
	"net/http"
	"log"
	"strings"
)
/*
将HTTP请求转发给数据服务
*/
func put(w http.ResponseWriter, r *http.Request) {
	//  从URL中解析<object_name>部分赋值给object
	object := strings.Split(r.URL.EscapedPath(),"/")[2]
	c , e := storeObject(r.Body,object)
	if e != nil {
		log.Println(e)
	}
	//  错误代码写入HTTP响应中
	w.WriteHeader(c)
}

