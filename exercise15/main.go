package main

import (
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {


		id, err := uuid.NewV4()

		logError(err)

		cookie := &http.Cookie{
			Name:     "session-id",
			Value:    id.String(),
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	})

	log.Println("Listening to 8080 ...")
	http.ListenAndServe(":8080", nil)
}

func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}
