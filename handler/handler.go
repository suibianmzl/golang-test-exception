package handler

import (
	"github.com/suibianmzl/test-exception/exception"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const PathPrefix = "/list/"

// 实现一个读取文件的 httpServer 处理器
// 假设访问 http://localhost:8888/list/abc.txt
func HandFileListing(writer http.ResponseWriter, request *http.Request) error {

	// 1. 如果 urlPath中不是以list开头，就返会自定义用户类型错误
	if strings.Index(request.URL.Path, PathPrefix) != 0 {
		return exception.MyCustomError("url path need startWith /list/")
	}
	//fmt.Println("path", request.URL.Path)    // /list/abc.txt
	path := request.URL.Path[len(PathPrefix):] // abc.txt 字符串切割，substring

	// 2. 打开文件
	file, err := os.Open(path)
	if err != nil {
		// 遇到错误直接返回，又错误统一处理器处理
		return err
	}
	defer file.Close()

	// 3. 读取文件到 byte[]
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// 4. 将byte[] all 写入相应流
	writer.Write(all)

	return nil
}

