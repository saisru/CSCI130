package main

import (
  "fmt"
  "io"
  "net/http"
)

// function that handles the website's back end
func thisLittleWebpage(res http.ResponseWriter, req *http.Request) {
  // for simplicity, the webpage is generated here
  page := `
    <!DOCTYPE html>
    <html>
      <head>
        <meta charset="utf-8">
        <title></title>
      </head>
      <body>
        <form method = "POST" enctype="multipart/form-data"> Input your name:
          <input type="file" name="name"><br>
          <input type="submit">
        </form>
      </body>
    </html>`

  io.WriteString(res, page) // write the page

	if req.Method == "POST" {
		_, src, err := req.FormFile("name")
		if err != nil {
			fmt.Println(err)
		}

		dst, err := src.Open()
		if err != nil {
			fmt.Println(err)
		}

		io.Copy(res, dst)
	}
}

func main() {
  http.HandleFunc("/", thisLittleWebpage)

  fmt.Println("server is now running...") // display when server is running
  http.ListenAndServe(":8080", nil) // set listener to port 8080 on localhost
}
