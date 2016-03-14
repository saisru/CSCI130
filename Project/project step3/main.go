

package main
import (
		"net/http"					
		"html/template" 				
		"io/ioutil"						
		"github.com/nu7hatch/gouuid"				
		"strings"						
 		"fmt" 									
)
func main() {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		name := req.FormValue("name")
		age := req.FormValue("age")
		cookie, err := req.Cookie("session-fino")
		if err != nil {
			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name:  "session-fino",
				Value: id.String() + "|" + name + age,
				// Secure: true,
				HttpOnly: true,
			}
			http.SetCookie(res, cookie)
		}
		err = tpl.Execute(res, nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}