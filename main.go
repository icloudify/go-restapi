package main

import (
	"fmt"
	"io"
	"time"

	//"github.com/ravindra031/go-restapi/gorilla"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat("#", 5), " Golang RESTAPI ", strings.Repeat("#", 5))
	/*router := httprouter.New()
	router.GET("/api/v1/go-version", custhttprouter.GoVersion)
	router.GET("/api/v1/show-file/:name", custhttprouter.GetFileContent)
	log.Fatal(http.ListenAndServe(":8080", router))*/
	//var r *mux.Router
	//r = mux.NewRouter()
	//r.Host("") CORS allow origin domain
	//1
	//r.HandleFunc("/articles/{category}/{id:[0-9]+}", gorilla.ArticalHandler)

	//2
	//r.Path("/articles/{category}/{id:[0-9]+}").HandlerFunc(gorilla.ArticalHandler)

	//3
	/*r.PathPrefix("/articles").Subrouter()
	r.HandleFunc("/{id}/setting",gorilla.SettingsHandler)
	r.HandleFunc("/{id}/details", gorilla.DetailHandler)*/

	//4
	/*r.StrictSlash(true)
	rr := &gorilla.Route{}
	r.Path("/articles/").Handler(rr)*/

	//5
	/*r.StrictSlash(true)
	r.HandleFunc("/articles", gorilla.QueryHandler)
	r.Queries("id", "category")*/

	/*mainLogicHandler := http.HandlerFunc(gomiddleware.MainLogic)
	Handler: gomiddleware.Middleware(mainLogicHandler)*/

	// RESTFul webservices
	webservice := new(restful.WebService)
	webservice.Route(webservice.GET("/ping").To(pingTime))
	restful.Add(webservice)

	srv := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	log.Fatal(srv.ListenAndServe())

}

func pingTime(request *restful.Request, response *restful.Response) {
	io.WriteString(response, fmt.Sprintf("%s", time.Now()))
}
