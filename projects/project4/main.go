package main

import (
	"net/http"
	"html/template"
	"log"
	"encoding/json"
	"encoding/base64"
	"strings"
)


type User struct {
	Name string
	Age string
}


func set_user(req *http.Request, user *User) string {
	user.Name = req.FormValue("user_name")
	user.Age = req.FormValue("user_age")

	bs, _ := json.Marshal(user)

	b64 := base64.URLEncoding.EncodeToString(bs)

	return b64
}

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
			HttpOnly: true,
		}
	}


	user := new(User)
	b64 := set_user(req, user)

	if req.FormValue("user_name") != "" {

		cookie_id := strings.Split(cookie.Value, "|")
		cookie.Value = cookie_id[0] + "|" + b64
	}

	http.SetCookie(res, cookie)

	tpl.Execute(res, nil)
}


func main() {

	http.ListenAndServe(":8080", nil)
}
