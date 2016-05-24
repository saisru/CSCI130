
package main

import (
	"fmt"
	"io"
	"net/http"
)

func somename(res http.ResponseWriter, req *http.Request) {
	val := req.FormValue("n")
	page := `
    <!DOCTYPE html>
    <html>
      <head>
        <meta charset="utf-8">
        <title></title>
      </head>
      <body>
        <form method="GET">
          <input type="text" name="n">
          <input type="submit">
        </form>
      </body>
    </html>`
	io.WriteString(res, page + val)
}

func main() {

	http.HandleFunc("/", somename)

	fmt.Println("server is now running...")
	http.ListenAndServe(":8080", nil)
}