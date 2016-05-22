package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	name := strings.Split(req.URL.Path, "/")
	nameOutput := strings.Join(name, "")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, nameOutput)
}