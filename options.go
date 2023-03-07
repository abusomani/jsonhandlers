package easyjson

import (
	"net/http"

	"github.com/abusomani/easyjson/handler"
)

// Option is a self-referential function to make palette modular and extensible.
// Read more about self-referential function here:
// https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
type Option func(*EasyJSON)

// WithHTTPRequestHandler is a self-referential function that takes in a http.ResponseWriter
// and a pointer to http.Request
// It sets the handler as HTTPRequestHandler on the EasyJSON receiver.
func WithHTTPRequestHandler(rw http.ResponseWriter, req *http.Request) Option {
	return func(e *EasyJSON) {
		h := handler.NewHTTPRequestHandler(rw, req)
		e.handler = h
	}
}

// WithHTTPResponseHandler is a self-referential function that takes in a pointer to http.Response
// It sets the handler as HTTPResponseHandler on the EasyJSON receiver.
func WithHTTPResponseHandler(res *http.Response) Option {
	return func(e *EasyJSON) {
		h := handler.NewHTTPResponseHandler(res)
		e.handler = h
	}
}

// WithFileHandler is a self-referential function that takes in a fileName
// It sets the handler as FileHandler on the EasyJSON receiver.
func WithFileHandler(fileName string) Option {
	return func(e *EasyJSON) {
		h := handler.NewFileHandler(fileName)
		e.handler = h
	}
}

// WithDefaults is a self-referential function.
// It sets the NoopHandler on the EasyJSON receiver.
func WithDefaults() Option {
	return func(e *EasyJSON) {
		h := handler.NewNoopHandler()
		e.handler = h
	}
}
