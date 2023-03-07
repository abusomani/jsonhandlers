package handler

import (
	"fmt"
	"io"
	"net/http"
)

// HTTPRequestHandler implements the Handler interface
// It takes in pointer to a http.Request and a http.ResponseWriter on which the operations have to be performed.
type HTTPRequestHandler struct {
	r  *http.Request
	rw http.ResponseWriter
}

// NewHTTPRequestHandler takes in http.ResponseWriter and pointer to a http.Request.
// It returns a HTTPRequestHandler instance.
func NewHTTPRequestHandler(rw http.ResponseWriter, req *http.Request) *HTTPRequestHandler {
	return &HTTPRequestHandler{
		r:  req,
		rw: rw,
	}
}

// Read function returns the bytes read from the given http.Request or an error in case something went wrong.
func (rh *HTTPRequestHandler) Read() ([]byte, error) {
	if rh.r == nil {
		return nil, fmt.Errorf("request is nil")
	}

	if rh.r.Body == nil {
		return nil, fmt.Errorf("request body is nil")
	}

	body, err := io.ReadAll(rh.r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading body: %s", err.Error())
	}

	return body, nil
}

// Write function writes the bytes of data using the given http.ResponseWriter.
// It returns an error in case something went wrong.
func (rh *HTTPRequestHandler) Write(input []byte) error {
	if rh.rw == nil {
		return fmt.Errorf("no writer configured")
	}
	_, err := rh.rw.Write(input)
	return err
}
