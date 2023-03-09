package jsonhandlers

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/abusomani/jsonhandlers/handler"
)

// JSONHandler exposes Marshal and Unmarshal methods.
type JSONHandler struct {
	// default handler is NoopHandler
	// It can take values of FileHandler, HTTPRequestHandler, HTTPResponseHandler
	handler handler.Handler
}

// New returns a new JSONHandler instance with variable options to set the required configurations.
func New(opts ...Option) *JSONHandler {
	jh := &JSONHandler{}

	// set the default configurations
	defaults := WithDefaults()
	defaults(jh)

	for _, opt := range opts {
		opt(jh)
	}
	return jh
}

func (jh *JSONHandler) SetOptions(opts ...Option) *JSONHandler {
	for _, opt := range opts {
		opt(jh)
	}
	return jh
}

// Marshal writes the JSON encoding of v in the configured source of handler.
//
// JSON cannot represent cyclic data structures and Marshal does not
// handle them. Passing cyclic structures to Marshal will result in
// an error.
func (jh *JSONHandler) Marshal(v any) error {
	if jh.handler == nil {
		return fmt.Errorf("handler not configured")
	}

	value, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("json marshal: %s", err.Error())
	}

	if err := jh.handler.Write(value); err != nil {
		return fmt.Errorf("error in write: %s", err.Error())
	}

	return nil
}

// Unmarshal parses the JSON-encoded data and stores the result in the value given in input as `v`.
func (jh *JSONHandler) Unmarshal(v any) error {
	if jh.handler == nil {
		return fmt.Errorf("handler not configured")
	}
	data, err := jh.handler.Read()
	if err != nil {
		return errors.New(err.Error())
	}
	return json.Unmarshal(data, v)
}
