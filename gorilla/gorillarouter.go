package gorilla

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ArticalHandler(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func SettingsHandler(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func DetailHandler(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "details ID is: %v\n", vars["id"])
}

type Route struct {
}

func (r *Route) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Strict is called \n")
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("QueryHandler called")
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", queryParams["category"])
	fmt.Fprintf(w, "ID is: %v\n", queryParams["id"])
}
