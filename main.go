package main

import (
	"net/http"
	"fmt"

	"github.com/zingp/gowebk8sdeploy/app01"
)

func main() {
	port := "0.0.0.0:8888:8888"
	http.HandleFunc("/hello", app01.Hello)
	http.HandleFunc("/index", app01.Index)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("http listen failed:", err)
		return
	}
	fmt.Printf("listen %s start...", port)
}