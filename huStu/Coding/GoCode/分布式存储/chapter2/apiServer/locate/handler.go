package locate

import (
	"net/http"
	"encoding/json"
	"strings"
)

func Handler(w http.ResponseWriter , r * http.Request){
	m := r.Method
	//  不是Get请求
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return 
	}
	info := Locate(strings.Split(r.URL.EscapedPath(),"/")[2])
	if len(info) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b , _ := json.Marshal(info)
	w.Write(b)
}

