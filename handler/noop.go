package handler

// Noop implements the Handler interface
// This is a default handler when no input is provided.
type Noop struct{}

// NewNoopHandler takes returns a Noop instance.
func NewNoopHandler() *Noop {
	return &Noop{}
}

// Read function returns the empty bytes and a nil error.
func (noop *Noop) Read() ([]byte, error) {
	return make([]byte, 0), nil
}

// Write operation returns a nil error.
func (noop *Noop) Write(d []byte) error {
	return nil
}
