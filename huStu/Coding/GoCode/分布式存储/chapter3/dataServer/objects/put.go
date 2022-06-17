package objects

import(
	"net/http"
	"os"
	"log"
	"io"
	"strings"
)

/*
改进:该实现如果用户的<object_name>中包含'/',则'/'后面会被丢弃,
	正确处理应该是返回给用户一个错误.
*/
func put(w http.ResponseWriter, r *http.Request) {
	//  创建上传的同名文件
	file , err := os.Create(os.Getenv("STORAGE_ROOT")+"/objects/" + 
							strings.Split(r.URL.EscapedPath(),"/")[2]) //  由于路径为/objects/<object_name>,分割以后为返回一个字符串数组 "","objects","object_name"
	if err != nil{
		log.Println(err)
		//  创建失败返回HTTP错误代码500
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}
	defer file.Close()
	//  数据拷贝到创建的file文件
	io.Copy(file,r.Body)
}