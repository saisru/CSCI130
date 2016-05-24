package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/plain")
		res.Write([]byte("SECURE\n"))
	})
	http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
}