package operations

import (
	"net/http"

	"github.com/abusomani/jsonhandlers"
)

type studentSearchRequest struct {
	Name string
}

type studentSearchResponse struct {
	Info student
}

func HandleHTTPRequest(students []student) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jh := jsonhandlers.New(jsonhandlers.WithHTTPRequestHandler(w, r))

		var reqBody studentSearchRequest
		err := jh.Unmarshal(&reqBody)
		if err != nil {
			errPayload := struct {
				StatusCode int
				Message    string
			}{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			}
			// request is bad
			jh.Marshal(errPayload)
			return
		}

		for _, student := range students {
			// student found
			if student.Name == reqBody.Name {
				// response has the right student info written
				jh.Marshal(studentSearchResponse{
					Info: student,
				})
				return
			}
		}

		errPayload := struct {
			StatusCode int
			Message    string
		}{
			StatusCode: http.StatusInternalServerError,
			Message:    "something went wrong",
		}
		// student not found
		jh.Marshal(errPayload)
	})
}

/*
  Sample request to be hit on the localhost server to test WithHTTPRequestHandler functionality.
  curl http://localhost:8080/search -d '{"Name": "Abhishek Somani"}'
*/
