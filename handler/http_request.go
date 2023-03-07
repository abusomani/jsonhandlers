package handler

import (
	"fmt"
	"io"
	"net/http"
)

// HttpRequestHandler implements the Handler interface
// It takes in pointer to a http.Request and a http.ResponseWriter on which the operations have to be performed.
type HttpRequestHandler struct {
	r  *http.Request
	rw http.ResponseWriter
}

// NewHttpRequestHandler takes in http.ResponseWriter and pointer to a http.Request.
// It returns a HttpRequestHandler instance.
func NewHttpRequestHandler(rw http.ResponseWriter, req *http.Request) *HttpRequestHandler {
	return &HttpRequestHandler{
		r:  req,
		rw: rw,
	}
}

// Read function returns the bytes read from the given http.Request or an error in case something went wrong.
func (rh *HttpRequestHandler) Read() ([]byte, error) {
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

// Write function writes the bytes of data using the given http.ResponseWriter and returns an error in case something went wrong.
func (rh *HttpRequestHandler) Write(input []byte) error {
	if rh.rw == nil {
		return fmt.Errorf("no writer configured")
	}
	_, err := rh.rw.Write(input)
	return err
}
