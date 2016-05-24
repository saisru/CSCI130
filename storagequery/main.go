
package storage

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	storageLog "google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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

	// Storing a dummy file using different delimiters (will be stored in different substrings)
	store(ctx, client, " folder/")
	store(ctx, client, " folderlike+")

	// These look like actual folders on storage
	printFolders(ctx, client, res, "/")

	// These do not look like folders on storage since the default delimeter is not used
	printFolders(ctx, client, res, "+")
}

func printFolders(ctx context.Context, client *storage.Client, res http.ResponseWriter, delimeter string) {
	fmt.Fprintf(res, "Delimeter ["+delimeter+"]\n")
	query := &storage.Query{
		Delimiter: delimeter,
	}
	objs, err := client.Bucket(BUCKET_NAME).List(ctx, query)
	logError(err)
	for _, subfolder := range objs.Prefixes {
		fmt.Fprintf(res, "Folder: "+subfolder+"\n")
	}
}

// Stores a dummy file in a loop by using the folder postfix string passed.
func store(ctx context.Context, client *storage.Client, folderPostfix string) {
	// Reading the file from disk
	reader, err := os.Open(FILE_NAME)
	logError(err)

	// Looping to create few number of folders
	for i := 0; i < 3; i++ {
		// Adding i index as a prefix to the path in which we want to create the files on storage.
		writer := client.Bucket(BUCKET_NAME).Object(strconv.Itoa(i) + folderPostfix + FILE_NAME).NewWriter(ctx)
		writer.ACL = []storage.ACLRule{{
			storage.AllUsers,
			storage.RoleReader}}
		io.Copy(writer, reader)
		writer.Close()
	}
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