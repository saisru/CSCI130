package main

import (
	"net/http"
	"io"
)

func main(){
	http.HandleFunc("/",func(res http.ResponseWriter,req *http.Request) {
		io.WriteString(res,"hello GO")
	})
	http.ListenAndServe(":9000",nil)
}