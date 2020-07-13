package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"strings"
	"net/http"
	"net/url"
	//"html/template"
	"time"
)

var port = ":1234"

type website struct {
	Title string
	Time time.Time
}

func main() {
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
	http.HandleFunc("/",RequestLog)
	log.Print("Listen Port",port)
	err := http.ListenAndServe(port,nil)

	if err != nil {
		log.Fatal("ListenAndServe: ",err)
	}
}

func RequestLog(w http.ResponseWriter, req *http.Request) {

	//teml,err := template.ParseFiles("???.html") //動的に使いたいときに使う
	//if err != nil {
	//	log.Print(err)
	//}

	//Log Check
	method := req.Method
	fmt.Println()
	log.Println("[method]" + method)
	for k,v := range req.Header {
		fmt.Print("[header]" + k)
		fmt.Println(": " + strings.Join(v, ","))
	}
	switch method {
	case "GET":
		req.ParseForm()

		for k, v := range req.Form {
			fmt.Print("[param]" + k)
			fmt.Println(": " + strings.Join(v, ","))
		}
	case "POST":
		defer req.Body.Close()
		body,err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Print(err)
		}
		decoded,err := url.QueryUnescape(string(body))
		if err != nil {
			log.Print(err)
		}
		fmt.Println("[Request Body]",decoded)
	default:
		log.Print("Method not allowed.\n")
	}

}
