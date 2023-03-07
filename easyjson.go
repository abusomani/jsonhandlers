package easyjson

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/abusomani/easyjson/handler"
)

// EasyJSON exposes Marshal and Unmarshal methods.
type EasyJSON struct {
	// default handler is NoopHandler
	// It can take values of FileHandler, HTTPRequestHandler, HTTPResponseHandler
	handler handler.Handler
}

// New returns a new EasyJSON instance with variable options to set the required configurations.
func New(opts ...Option) *EasyJSON {
	e := &EasyJSON{}

	// set the default configurations
	defaults := WithDefaults()
	defaults(e)

	for _, opt := range opts {
		opt(e)
	}
	return e
}

// Marshal writes the JSON encoding of v in the configured source of handler.
//
// JSON cannot represent cyclic data structures and Marshal does not
// handle them. Passing cyclic structures to Marshal will result in
// an error.
func (e *EasyJSON) Marshal(v any) error {
	if e.handler == nil {
		return fmt.Errorf("handler not configured")
	}

	value, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("json marshal: %s", err.Error())
	}

	if err := e.handler.Write(value); err != nil {
		return fmt.Errorf("error in write: %s", err.Error())
	}

	return nil
}

// Unmarshal parses the JSON-encoded data and stores the result in the value given in input as `v`.
func (e *EasyJSON) Unmarshal(v any) error {
	if e.handler == nil {
		return fmt.Errorf("handler not configured")
	}
	data, err := e.handler.Read()
	if err != nil {
		return errors.New(err.Error())
	}
	return json.Unmarshal(data, v)
}
