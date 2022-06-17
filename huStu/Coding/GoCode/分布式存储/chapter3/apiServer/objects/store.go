package objects

import (
	"io"
	"net/http"
)

//  将HTTP正文数据保存在stream中
func storeObject(r io.Reader,object string) (int ,error) {
	//  生成objectstream.PutStream
	stream ,e := putSreanm(object)
	if e != nil {
		return http.StatusServiceUnavailable,e
	}
	//  将HTTP的正文数据写入流中
	io.Copy(stream,r)
	//  关闭流
	e = stream.Close()
	if e != nil {
		return http.StatusInternalServerError,e
	}
	return http.StatusOK,nil
}