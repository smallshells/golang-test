package main

import (
	"fmt"
	"net/http"
)

// HTTP Server

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1 style='color:green'>hello world!</h1>")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server failed，err：", err)
		return
	}
}
