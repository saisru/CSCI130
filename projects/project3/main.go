package main

import (
	"net/http"
	"html/template"
	"log"
	"strings"
)

func serve_the_webpage(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("appform.html")
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
	}

	if req.FormValue("user_name") != "" && !strings.Contains(cookie.Value, "user_name") {
		cookie.Value = cookie.Value +
		`name=` + req.FormValue("user_name") +
		`age=` + req.FormValue("user_age")
	}

	http.SetCookie(res, cookie)

	tpl.Execute(res, nil)
}


func main() {
	http.ListenAndServe(":8080", nil)
}