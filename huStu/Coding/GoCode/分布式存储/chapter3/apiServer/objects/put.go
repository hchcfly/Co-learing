package objects

import(
	"net/http"
	"log"
	"strings"
	"net/url"
	"../../../lib/es"
	"../../../lib/utils"
)
/*
将HTTP请求转发给数据服务
*/
func put(w http.ResponseWriter, r *http.Request) {
	//  从http请求头部获取对象的hash值
	hash := utils.GetHashFromHeader(r.Header)
	if hash == "" {
		log.Println("missing object hash in digest header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//  根据对象的hash值存储对象
	c , e := storeObject(r.Body,url.PathEscape(hash))
	if e != nil {
		log.Println(e)
		w.WriteHeader(c)
		return
	}
	if c != http.StatusOK {
		w.WriteHeader(c)
		return
	}
	//  从URL中获取对象名称
	name := strings.Split(r.URL.EscapedPath(),"/")[2]
	//  从URL中获取对象的size
	size := utils.GetSizeFromHeader(r.Header)
	//  向es中添加新版本对象的metadata
	e = es.AddVersion(name,hash,size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
}



