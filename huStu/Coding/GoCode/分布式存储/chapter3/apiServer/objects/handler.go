package objects

import "net/http"
import "fmt"

func Handler(w http.ResponseWriter , r * http.Request){
	m := r.Method
	fmt.Println("method = ",m)
	if m == http.MethodPut {
		put(w,r)
	}
	if m == http.MethodGet {
		get(w,r)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}