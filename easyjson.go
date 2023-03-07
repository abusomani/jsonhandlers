package easyjson

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/abusomani/easyjson/handler"
)

// EasyJson exposes Marshal and Unmarshal method
type EasyJson struct {
	// default handler is NoopHandler
	// It can take values of FileHandler, HttpRequestHandler, HttpResponseHandler
	handler handler.Handler
}

// New returns a new EasyJson instance with variable options to set the required configurations.
func New(opts ...Option) *EasyJson {
	e := &EasyJson{}

	// set the default configurations
	defaults := WithDefaults()
	defaults(e)

	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e *EasyJson) Marshal(v any) error {
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

func (e *EasyJson) Unmarshal(v any) error {
	if e.handler == nil {
		return fmt.Errorf("handler not configured")
	}
	data, err := e.handler.Read()
	if err != nil {
		return errors.New(err.Error())
	}
	return json.Unmarshal(data, v)
}
