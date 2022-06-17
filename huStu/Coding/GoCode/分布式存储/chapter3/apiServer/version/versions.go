package versions

import (
	"net/http"
	"strings"
	"log"
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
		metas,e := es.SearchAllVersions(name,from,size)
		if e != nil {
			log.Println(e)
		}
	}


}