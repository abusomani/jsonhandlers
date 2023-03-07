package easyjson

import "github.com/abusomani/easyjson/handler"

type Option func(*EasyJson)

func WithHandler(h handler.Handler) Option {
	return func(e *EasyJson) {
		e.handler = h
	}
}
