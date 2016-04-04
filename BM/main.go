package main

import (
	"github.com/nu7hatch/gouuid"
	"html/template"
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha256"
	"io"
	"strings"
)

type User struct {
	Name string
	Age string
	}

func main() {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		name := req.FormValue("name")
		age := req.FormValue("age")
		info := stuff(name, age)
		code := getCode(info)
		cookie, err := req.Cookie("session-fino")
		if err != nil {
			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name:  "session-fino",
				Value: id.String() + "|" + name + age + "|" + info + "|" + code,
				// Secure: true,
				HttpOnly: true,
			}
			http.SetCookie(res, cookie)

		}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	cookVal := strings.Split(cookie.Value, "|")
	Uuid := cookVal[0]
	CodeUuid := cookVal[1]
	if (getCode(Uuid) == CodeUuid) {
		fmt.Fprintf(res, "Everything's fine\n")
	} else {
		fmt.Fprintf(res,"No good \n")
	}


		err = tpl.Execute(res, nil)
		if err != nil {
			log.Fatalln(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func stuff(name string, age string) string{
person := User{
			Name: name, 
			Age: age,
		}
	
		b, err := json.Marshal(person)
		if err != nil {
			fmt.Printf("error: " , err)}
		encode := base64.StdEncoding.EncodeToString(b)
		return encode
