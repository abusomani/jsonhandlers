package operations

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abusomani/jsonhandlers"
)

type user struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
	Gender    string
	Email     string
}

type getUsersResponse struct {
	Users []user
}

func HandleHTTPResponse() {
	resp, err := http.Get("https://dummyjson.com/users")
	if err != nil {
		log.Fatalf("unable to make the get request %s", err.Error())
	}
	jh := jsonhandlers.New(jsonhandlers.WithHTTPResponseHandler(resp))

	var userResp getUsersResponse
	jh.Unmarshal(&userResp)
	fmt.Printf("response is %+v\n", userResp)
}
