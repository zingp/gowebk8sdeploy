package app01

import(
	"net/http"
	"fmt"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "hello web")
	return
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle index")
	fmt.Fprintf(w, "index web")
	
}

