package objects

import (
	"net/http"
	"strings"
	"log"
	"../../../lib/es"
)

//  删除对应的对象.
//  并不是真正的删除,实际还存储在数据存储层.
func del(w http.ResponseWriter,r *http.Request) {
	//  获取object_name
	name := strings.Split(r.URL.EscapedPath(),"/")[2]
	//  获取最新的版本数据metadata
	version ,e := es.SearchLatestVersion(name)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}
	//  插入删除标记
	//  对象名称:版本号:对象大小:哈希值
	e = es.PutMetadata(name,version.Version+1,0,"")
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}
}
