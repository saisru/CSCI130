package main

import (
	"net/http"
	"html/template"
	"log"
	"fmt"
)


func cookie_id(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("test.html")
	if err != nil {
		log.Fatalln(err)
	}
	cookie, err := req.Cookie("session-information")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-information",
			Value: id.String(),
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}

	fmt.Println(cookie_id)
}


func main() {
	http.HandleFunc("/", cookie_id)
	http.ListenAndServe(":8080", nil)
}
