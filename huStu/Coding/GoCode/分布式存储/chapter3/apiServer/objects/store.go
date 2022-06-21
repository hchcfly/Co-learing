package objects

import (
	"io"
	"net/http"
)

//  将HTTP正文数据保存在stream中
//  size用来确定临时对象的大小
func storeObject(r io.Reader,hash string,size int64) (int ,error) {
	//定位对象的散列值是否存在
	if local.Exist(url.PathEscape(hash)) {
		return http.StatusOK,nil
	}
	
	//  生成objectstream.PutStream
	stream ,e := putSreanm(url.PathEscape(hash),size)
	if e != nil {
		return http.StatusServiceUnavailable,e
	}
	//将数据读入临时对象中
	reader := io.TeeReader(r,stream)
	d := utils.CalculateHash(reader)
	if d != hash {
		//  删除临时对象
		stream.Commit(false)
		return http.StatusBadRequest,fmt.Errorf("object hash mismatch, calculated=%s, requested=%s", d, hash)
	}
	// //  将HTTP的正文数据写入流中
	// io.Copy(stream,r)
	// //  关闭流
	// e = stream.Close()
	// if e != nil {
	// 	return http.StatusInternalServerError,e
	// }
	//  将临时对象转正并返回
	stream.Commit(true)
	return http.StatusOK,nil
}