package main

import (
	"github.com/suibianmzl/test-exception/exception"
	"github.com/suibianmzl/test-exception/handler"
	"net/http"
)

func main()  {
	// 1. 注册处理 handler.PathPrefix 开头的业务逻辑处理器
	http.HandleFunc(handler.PathPrefix, exception.ErrWrapper(handler.HandFileListing))

	// 2. 启动 httpServer, 监听端口
	err := http.ListenAndServe("127.0.0.1:8888", nil)

	// 3.如果启动失败，则直接抛出错误
	if err != nil{
		panic(err)
	}
}
