package main

import (
	"net/http"
	"io"
	"log"
)

func main(){
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.HandleFunc("/",func(res http.ResponseWriter,req *http.Request){
		key := "name"
		val := req.FormValue(key)
		log.Println(val)
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(res, `<form method="POST">
		 <label for="Name">Your Name is => </label>
		 <input type="text" name="name">
		 <input type="submit" value="Click here to display name below">
		</form>`)
		if(val != ""){
			io.WriteString(res,"You Entered your name as " + val);
		}

	})
	http.ListenAndServe(":8080",nil)
}
