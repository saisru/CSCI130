package main

import (
	"html/template"
	"net/http"
)

func main() {
	tpl, err := template.ParseFiles("webapp.html")


	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		err = tpl.Execute(res, nil)

	})
	http.ListenAndServe(":8080", nil)
}
