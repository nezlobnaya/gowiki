package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//An http.Request is a data structure that represents the client HTTP request. r.URL.Path is the path component of the request URL.
	//The trailing [1:] means "create a sub-slice of Path from the 1st character to the end." This drops the leading "/" from the path name.
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

const appPort = "8081"

func main() {
	//The main function begins with a call to http.HandleFunc,
	//which tells the http package to handle all requests to the web root ("/") with handler.
	http.HandleFunc("/", handler)
	log.Println("Listening on port", appPort)
	//ListenAndServe always returns an error, since it only returns when an unexpected error occurs.
	//In order to log that error we wrap the function call with log.Fatal.
	log.Fatal(http.ListenAndServe(":8081", nil))
}
