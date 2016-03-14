
package main

import (
	"fmt"
	"strings"
	"net/http"
)

func serve_the_webpage(res http.ResponseWriter, req *http.Request) {
	fs := strings.Split(req.URL.Path, "/")
	fmt.Fprint(res, fs[1]) 
}

func main() {
	http.HandleFunc("/", serve_the_webpage)

	http.ListenAndServe(":8080", nil)
}