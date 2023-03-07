package easyjson

import (
	"net/http"

	"github.com/abusomani/easyjson/handler"
)

// Option is a self-referential function to make palette modular and extensible.
// Read more about self-referential function here:
// https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
type Option func(*EasyJson)

// WithHttpRequestHandler is a self-referential function that takes in a http.ResponseWriter and a pointer to http.Request
// It sets the handler as HttpRequestHandler on the EasyJson receiver.
func WithHttpRequestHandler(rw http.ResponseWriter, req *http.Request) Option {
	return func(e *EasyJson) {
		h := handler.NewHttpRequestHandler(rw, req)
		e.handler = h
	}
}

// WithHttpResponseHandler is a self-referential function that takes in a pointer to http.Response
// It sets the handler as HttpResponseHandler on the EasyJson receiver.
func WithHttpResponseHandler(res *http.Response) Option {
	return func(e *EasyJson) {
		h := handler.NewHttpResponseHandler(res)
		e.handler = h
	}
}

// WithFileHandler is a self-referential function that takes in a fileName
// It sets the handler as FileHandler on the EasyJson receiver.
func WithFileHandler(fileName string) Option {
	return func(e *EasyJson) {
		h := handler.NewFileHandler(fileName)
		e.handler = h
	}
}

// WithDefaults is a self-referential function.
// It sets the NoopHandler on the EasyJson receiver.
func WithDefaults() Option {
	return func(e *EasyJson) {
		h := handler.NewNoopHandler()
		e.handler = h
	}
}
