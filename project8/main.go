package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
	"text/template"
	"crypto/hmac"
	"io"
	"crypto/sha256"
)

type User struct {
	Age  string
	Name string
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("projectlogin"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func startpage(res http.ResponseWriter, req *http.Request) {
	var err error
	temp := template.Must(template.ParseFiles("index.html"))
	name := req.FormValue("name")
	age := req.FormValue("age")
	UserDetails := User{
		Age:  age,
		Name: name,
	}

	bs, err := json.Marshal(UserDetails)
	if err != nil {
		fmt.Println(err)
	}

	json := base64.StdEncoding.EncodeToString(bs)

	cookie, err := req.Cookie("session-fino")
	if err == http.ErrNoCookie {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String(),
			//Secure : true,
			HttpOnly: true,
		}
	} else {
		if req.Method == "POST" {
			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name:  "session-fino",
				Value: id.String() + name + age + json,
				//Secure : true,
				HttpOnly: true,
			}
		}
	}
	http.SetCookie(res, cookie)
	err = temp.Execute(res, nil)
	if err != nil {
		panic(err)
	}
}
func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", startpage)
	http.ListenAndServe(":8080", nil)
}
