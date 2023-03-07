package handler

type Handler interface {
	Parse([]byte, any)
}
