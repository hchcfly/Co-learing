package objects

import "net/http"
import "fmt"

func Handler(w http.ResponseWriter , r * http.Request){
	m := r.Method
	fmt.Println("method = ",m)
	if m == http.MethodPut {
		put(w,r)
		return 
	}
	if m == http.MethodGet {
		get(w,r)
		return
	}
	if m == http.MethodDelete {
		del(w,r)
		return 
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}