package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	
)

func handler_simple(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi\n")
}

func handler_readfile(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("./web.go")
	fmt.Fprintf(w, string(data))
}

func main() {
	http.HandleFunc("/", handler_simple)
	http.HandleFunc("/file", handler_readfile)
	http.ListenAndServe(":8080", nil)
}
