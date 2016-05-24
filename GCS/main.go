
package storage

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	storageLog "google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"io"
	"log"
	"net/http"
	"os"
)

const BUCKET_NAME = "csci130-goapp.appspot.com"
const FILE_NAME = "randomtext.txt"

func init() {
	http.HandleFunc("/", handler)
}

func handler(res http.ResponseWriter, req *http.Request) {

	// Saving the file
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	logStorageError(ctx, "Could not create a new client", err)
	defer client.Close()

	writer := client.Bucket(BUCKET_NAME).Object(FILE_NAME).NewWriter(ctx)
	writer.ACL = []storage.ACLRule{{
		storage.AllUsers,
		storage.RoleReader}}

	// Reading the file from disk
	reader, err := os.Open(FILE_NAME)
	logError(err)
	io.Copy(writer, reader)
	writer.Close()

	// Reading the file
	rd, err := client.Bucket(BUCKET_NAME).Object(FILE_NAME).NewReader(ctx)
	logError(err)
	io.Copy(res, rd)
	defer rd.Close()
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