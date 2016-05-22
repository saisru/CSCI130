
package main

import (
	"net/http"
	"fmt"
	"github.com/nu7hatch/gouuid"

)


func cookieHandler(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("session-id")
	if err == http.ErrNoCookie{
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-id",
			Value: id.String(),
			HttpOnly:true,

		}
	}
	http.SetCookie(res, cookie)
	fmt.Fprintf(res, "<h1>Check counter-cookie by developer tool to assert sessionID is:%s</h1>",
		cookie.Value)
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", cookieHandler)
	http.ListenAndServe(":8080", nil)
}