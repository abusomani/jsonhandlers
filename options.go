package jsonhandlers

import (
	"net/http"

	"github.com/abusomani/jsonhandlers/handler"
)

// Option is a self-referential function to make palette modular and extensible.
// Read more about self-referential function here:
// https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
type Option func(*JSONHandler)

// WithHTTPRequestHandler is a self-referential function that takes in a http.ResponseWriter
// and a pointer to http.Request
// It sets the handler as HTTPRequestHandler on the JSONHandler receiver.
func WithHTTPRequestHandler(rw http.ResponseWriter, req *http.Request) Option {
	return func(jh *JSONHandler) {
		h := handler.NewHTTPRequestHandler(rw, req)
		jh.handler = h
	}
}

// WithHTTPResponseHandler is a self-referential function that takes in a pointer to http.Response
// It sets the handler as HTTPResponseHandler on the JSONHandler receiver.
func WithHTTPResponseHandler(res *http.Response) Option {
	return func(jh *JSONHandler) {
		h := handler.NewHTTPResponseHandler(res)
		jh.handler = h
	}
}

// WithFileHandler is a self-referential function that takes in a fileName
// It sets the handler as FileHandler on the JSONHandler receiver.
func WithFileHandler(fileName string) Option {
	return func(jh *JSONHandler) {
		h := handler.NewFileHandler(fileName)
		jh.handler = h
	}
}

// WithDefaults is a self-referential function.
// It sets the NoopHandler on the JSONHandler receiver.
func WithDefaults() Option {
	return func(jh *JSONHandler) {
		h := handler.NewNoopHandler()
		jh.handler = h
	}
}
