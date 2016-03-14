package main 

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
)

func EnterName(res http.ResponseWriter, req *http.Request){
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(res, nil)
	fmt.Fprintf(res, "%v", req.FormValue("name"))
}

func main() {
	http.HandleFunc("/", EnterName)
	http.ListenAndServe(":8080", nil)
}