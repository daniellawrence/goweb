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
func handler_bigresponse(w http.ResponseWriter, r *http.Request) {
	bytes := make([]byte, 5000)
	for i, _ := range bytes {
		bytes[i] = 'a'
	}
	fmt.Fprintf(w, string(bytes))
}

func main() {
	http.HandleFunc("/", handler_simple)
	http.HandleFunc("/file", handler_readfile)
	http.HandleFunc("/big", handler_bigresponse)
	http.ListenAndServe(":8080", nil)
}
