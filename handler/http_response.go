package handler

import (
	"fmt"
	"io"
	"net/http"
)

// HTTPResponseHandler implements the Handler interface
// It takes in pointer to a http.Response on which the operations have to be performed.
type HTTPResponseHandler struct {
	r *http.Response
}

// NewHTTPResponseHandler takes in a pointer to http.Response.
// It returns a HTTPResponseHandler instance.
func NewHTTPResponseHandler(res *http.Response) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		r: res,
	}
}

// Read function returns the bytes read from the given http.Response or an error in case something went wrong.
func (rh *HTTPResponseHandler) Read() ([]byte, error) {
	if rh.r == nil {
		return nil, fmt.Errorf("response is nil")
	}
	if rh.r.Body == nil {
		return nil, fmt.Errorf("response body is nil")
	}
	defer func() {
		if err := rh.r.Body.Close(); err != nil {
			fmt.Printf("error closing : %s", err.Error())
		}
	}()
	body, err := io.ReadAll(rh.r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading body: %s", err.Error())
	}
	return body, nil
}

// Write operation is not allowed on a http.Response.
func (rh *HTTPResponseHandler) Write(input []byte) error {
	return fmt.Errorf("not allowed")
}
