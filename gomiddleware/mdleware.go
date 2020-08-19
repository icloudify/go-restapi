package gomiddleware

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	p := fmt.Println
	numGenerator := generator()
	for i := 0; i < 5; i++ {
		p("for loop ..", i)
		p(numGenerator(), "\t")
	}
}

func generator() func() int {
	var i = 0
	fmt.Println("Calling func ..", i)

	return func() int {
		fmt.Println("in anonymous func ...", i)
		i++
		return i
	}
}

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Executing middleware before request phase!")
		// Pass control back to the handler
		handler.ServeHTTP(writer, request)
		fmt.Println("Executing middleware after response phase!")
	})
}

func MainLogic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helllo middleware ...")
	io.WriteString(w, "helllo middleware ...")
}