package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoopRead(t *testing.T) {
	nh := NewNoopHandler()
	res, err := nh.Read()
	assert.Equal(t, make([]byte, 0), res)
	assert.Equal(t, nil, err)
}

func TestNoopWrite(t *testing.T) {
	nh := NewNoopHandler()
	err := nh.Write([]byte(""))
	assert.Equal(t, nil, err)
}
