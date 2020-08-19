package custhttprouter

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GoVersion(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("Version called")
	io.WriteString(w, "1.1")
}

func GetFileContent(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("File called called")
	file, error := os.Open("custhttprouter/" + params.ByName("name"))

	if error != nil {
		io.WriteString(w, error.Error())
	} else {
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Data as string: %s\n", data)
		io.WriteString(w, string(data))
	}
}
