package handler

import (
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

type mockReader struct {
	reader io.Reader
}

func newMockReader(r io.Reader) *mockReader {
	return &mockReader{
		reader: r,
	}
}

func (mr *mockReader) Read(p []byte) (int, error) {
	return mr.reader.Read(p)
}

func (mr *mockReader) Close() error {
	return errors.New("test error")
}

type noopWriter struct {
	throwsError bool
}

func newNoopWriter(te bool) *noopWriter {
	return &noopWriter{
		throwsError: te,
	}
}

func (nw *noopWriter) Write(data []byte) (n int, err error) {
	if nw.throwsError {
		return 0, errors.New("test error")
	}
	return 0, nil
}

func (nw *noopWriter) WriteHeader(statusCode int) {
}

func (nw *noopWriter) Header() http.Header {
	return http.Header{}
}

func TestNoopWriter(t *testing.T) {
	t.Run("testing error case of Write", func(t *testing.T) {
		nw := newNoopWriter(true)
		res, err := nw.Write(make([]byte, 0))
		assert.Equal(t, 0, res)
		assert.Equal(t, errors.New("test error"), err)
	})
	t.Run("testing success case of Write", func(t *testing.T) {
		nw := newNoopWriter(false)
		res, err := nw.Write(make([]byte, 0))
		assert.Equal(t, 0, res)
		assert.Equal(t, nil, err)
	})
	t.Run("testing WriteHeader", func(t *testing.T) {
		nw := newNoopWriter(false)
		nw.WriteHeader(0)
	})
	t.Run("testing Header", func(t *testing.T) {
		nw := newNoopWriter(false)
		h := nw.Header()
		assert.Equal(t, http.Header{}, h)
	})
}
