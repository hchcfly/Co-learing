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
	stream , e := getStream(object)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return 
	}
	io.Copy(w,stream)
}