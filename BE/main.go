package main

import (
	"fmt"
	"net/http"
)

func servePage(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res,
	`<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>TLS Web Page</title>
			</head>
			<body>
				<h1>Hello TLS</h1>
			</body>
			</html>`)
}

func redirectTLS(res http.ResponseWriter, req *http.Request){
	http.Redirect(res, req, "https://127.0.0.1:10443/"+req.RequestURI, http.StatusMovedPermanently)
}

func main(){
	http.HandleFunc("/", servePage)
	//key and cert Im not going to upload to Github
	go http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", "nil")

}