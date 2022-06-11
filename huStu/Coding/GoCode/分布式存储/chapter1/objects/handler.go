package objects

import "net/http"

func Handler(w http.ResponseWriter , r * http.Request){
	m := r.Method
	if m == http.MethodPut {

	}
	if m == http.MethodGet {

	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}