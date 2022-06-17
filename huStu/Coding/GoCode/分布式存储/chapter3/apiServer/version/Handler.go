package versions

import (
	"net/http"
	"strings"
	"log"
	"encoding/json"
)
//  对应版本的处理
func Handler(w http.ResponserWriter,r *http.Request) {
	m := r.Method
	//  当不为GET方法
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//  用来在es中分页查询
	from := 0
	size := 1000
	name := strings.Split(r.URL.EscapedPath(),"/")[2]
	for {
		metas,e := es.SearchAllVersions(name,from,size) // metadata : objects
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}
		for i := range metas {
			//  将metas[i]处理为json编码
			b,_ := json.Marshal(metas[i])
			//  写入HTTP响应正文
			w.Write(b)
			//  为甚要加\n?
			w.Write([]byte("\n"))
		}
		//  es中没有数据了
		if len(metas) != size {
			return 
		}
		//  读取下一页数据
		from += size
	}
}