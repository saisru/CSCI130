
package main

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	storageLog "google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"html/template"
	"io"
	"log"
	"net/http"
)

const BUCKET_NAME = "csci130-goapp.appspot.com"

func init() {
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/show", showHandler)
}

// Hanlder to upload user's images
func userHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		file, header, err := req.FormFile("image")
		logError(err)
		userName := req.FormValue("userName")
		saveFile(req, userName, header.Filename, file)
		http.Redirect(res, req, "/show?userName="+userName, http.StatusFound)
		return
	}

	//Parsing the template
	tpl := template.Must(template.ParseFiles("user.html"))
	err := tpl.Execute(res, nil)
	logError(err)
}

func saveFile(req *http.Request, userName string, fileName string, file io.Reader) {
	fileName = userName + "/" + fileName

	// Creating new context and client.
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	logStorageError(ctx, "Could not create a new client", err)
	defer client.Close()

	writer := client.Bucket(BUCKET_NAME).Object(fileName).NewWriter(ctx)
	writer.ACL = []storage.ACLRule{{
		storage.AllUsers,
		storage.RoleReader}}

	// Reading the file from disk
	io.Copy(writer, file)
	writer.Close()
}

// Handler for showing the images
func showHandler(res http.ResponseWriter, req *http.Request) {

	// Creating new context and client.
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	logStorageError(ctx, "Could not create a new client", err)
	defer client.Close()

	//Parsing the template
	tpl := template.Must(template.ParseFiles("index.html"))
	err = tpl.Execute(res, getPhotoNames(ctx, client, getUserName(req)))
	logError(err)
}

// Returns user's name from request parameters.
func getUserName(req *http.Request) string {
	return req.FormValue("userName")
}

// Returns the name of the photos uploaded in google storage
func getPhotoNames(ctx context.Context, client *storage.Client, userName string) []string {

	query := &storage.Query{
		Delimiter: "/",
		Prefix:    userName + "/",
	}
	objs, err := client.Bucket(BUCKET_NAME).List(ctx, query)
	logError(err)

	var names []string
	for _, result := range objs.Results {
		names = append(names, result.Name)
	}
	return names
}

// Logs the error given into log
func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Logs the error given into storage log
func logStorageError(ctx context.Context, errMessage string, err error) {
	if err != nil {
		storageLog.Errorf(ctx, errMessage, err)
	}
}