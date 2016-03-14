package main

import(
	"net/http"
	"html/template"
)

func main(){

	tpl, err := template.ParseFiles("Index.html")

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		err = tpl.Execute(res, nil)
	})

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}