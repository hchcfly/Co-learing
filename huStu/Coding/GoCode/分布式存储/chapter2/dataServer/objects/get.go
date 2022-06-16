package objects

import(
	"net/http"
	"os"
	"log"
	"io"
	"strings"
)

func get(w http.ResponseWriter,r *http.Request) {
	//  获取要get的文件名称
	file ,err := os.Open(os.Getenv("STORAGE_ROOT") + 
						"/objects/" + strings.Split(r.URL.EscapedPath(),"/")[2])
	if err != nil {
		log.Println(err)
		//  打开文件出错将错误404返回给客户端
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer file.Close()
	io.Copy(w,file)
}