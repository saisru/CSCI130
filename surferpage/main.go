package main

import (
	"html/template"

	"log"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {


	tpl, err := template.ParseFiles("surferpage.html")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	http.Handle("/pic/", http.StripPrefix("/pic", http.FileServer(http.Dir("pic"))))
	http.ListenAndServe(":8080", nil)
}