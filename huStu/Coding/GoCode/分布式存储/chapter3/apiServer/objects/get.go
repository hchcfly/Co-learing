package objects

import(
	"net/http"
	"log"
	"io"
	"strings"
)

func get(w http.ResponseWriter,r *http.Request) {
	//  从URL中解析<object_name>部分赋值给object
	object := strings.Split(r.URL.EscapedPath(),"/")[2]
	//  查询版本号数组
	versionId := r.URL.Query()["version"]
	version := 0
	var e error
	if len(versionId) != 0{
		//  获取对象对应版本号
		version,e = strconv.Atoi(versionId[0])
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	//  通过名称和版本号查询元数据
	meta,e := es.GetMetadata(name,version)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if meta.Hash == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//  获取对象的散列值
	object := url.PathEscape(meta.Hash)
	//  获取对应stram采用hash value
	stream,e := getStream(object)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return 
	}
	io.Copy(w,stream)
}