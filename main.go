package main

import (
	"fmt"
	"log"
	"io/util"
	"strings"
	"net/http"
	"net/url"
)

var port = ":1234"


func main() {
	http.HandleFunc("/hello",HelloServer)
	http.HandleFunc("/",HelloServer)
	log.Print("Listen Port",port)
	err := http.ListenAndServe(port,nil)

	if err != nil {
		log.Fatal("ListenAndServe: ",err)
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {

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
			log.Fatal(err)
		}
		decoded,error := url.QueryUnescape(string(body))
		if error != nil {
			log.Fatal(error)
		}
		fmt.Println("[Request Body]",decoded)

	default:
		log.Print("Method not allowed.\n")
	}
}
