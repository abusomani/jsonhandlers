package easyjson

import (
	"github.com/abusomani/easyjson/handler"
)

type EasyJson struct {
	handler handler.Handler
}

func New(opts ...Option) *EasyJson {
	e := &EasyJson{}

	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e *EasyJson) Parse(data []byte, v any) {
	e.handler.Parse(data, v)
}
