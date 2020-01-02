package exception

import (
	"log"
	"net/http"
	"os"
)

// 定义一个 function 的类型 type， 返回值是error
type addHandler func(writer http.ResponseWriter, request *http.Request) error

// 输入 appHandler 是一个函数, 输出也是一个函数 - 函数式编程
func ErrWrapper(handler addHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 1.处理业务逻辑
		err := handler(writer, request)
		if err != nil{
			log.Printf("error occured, %s", err)

			// 2.处理可以抛给用户的错误
			if err, ok := err.(UserError); ok {
				// 将错误写回到 http.ResponseWriter
				http.Error(writer, err.Message(), http.StatusBadRequest)
			}

			// 3.处理不可以抛给用户的错误
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)// 浏览器: Not Found
		}
	}
}
