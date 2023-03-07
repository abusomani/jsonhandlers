package handler

type Noop struct{}

func NewNoopHandler() *Noop {
	return &Noop{}
}
