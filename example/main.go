package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abusomani/jsonhandlers/example/operations"
)

func main() {
	operations.HandleHTTPResponse()
	fmt.Println()
	sts := operations.GetStudentsFromFile()
	fmt.Println()
	http.Handle("/search", operations.HandleHTTPRequest(sts))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("something went wrong while serving the http server")
	}
}
