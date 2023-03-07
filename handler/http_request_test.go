package handler

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpRequestRead(t *testing.T) {
	type Args struct {
		req *http.Request
		rw  http.ResponseWriter
	}
	tests := []struct {
		name    string
		wantErr bool
		args    Args
		err     error
		wantRes bool
		res     []byte
	}{
		{
			name:    "should return error if the request is nil",
			wantErr: true,
			err:     errors.New("request is nil"),
			args: Args{
				req: nil,
			},
			wantRes: false,
			res:     nil,
		},
		{
			name:    "should return error if the request body is nil",
			wantErr: true,
			err:     errors.New("request body is nil"),
			args: Args{
				req: &http.Request{},
			},
			wantRes: false,
			res:     nil,
		},
		{
			name:    "should return error if the reading of request body fails",
			wantErr: true,
			err:     errors.New("error while reading body: test error"),
			args: Args{
				req: httptest.NewRequest(http.MethodGet, "/something", errReader(0)),
			},
			wantRes: false,
			res:     nil,
		},
		{
			name:    "should return the bytes read from the request",
			wantErr: false,
			err:     nil,
			args: Args{
				req: httptest.NewRequest(http.MethodGet, "/something", ioutil.NopCloser(bytes.NewReader([]byte(`{"country": "brazil"}`)))),
			},
			wantRes: true,
			res:     []byte(`{"country": "brazil"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh := NewHTTPRequestHandler(tt.args.rw, tt.args.req)
			res, err := rh.Read()
			if tt.wantErr {
				assert.Equal(t, tt.err, err)
			}
			if tt.wantRes {
				assert.Equal(t, tt.res, res)
			}
		})
	}
}

func TestHttpRequestWrite(t *testing.T) {
	type Args struct {
		rw   http.ResponseWriter
		data []byte
	}
	tests := []struct {
		name    string
		wantErr bool
		args    Args
		err     error
	}{
		{
			name:    "should return error if the responseWriter is nil",
			wantErr: true,
			err:     errors.New("no writer configured"),
			args: Args{
				data: nil,
			},
		},
		{
			name:    "should return error in case write operation fails",
			wantErr: true,
			err:     errors.New("test error"),
			args: Args{
				data: []byte(`{"country": "brazil"}`),
				rw:   newNoopWriter(true),
			},
		},
		{
			name:    "should return the nil error in case of successful write operation",
			wantErr: false,
			err:     nil,
			args: Args{
				data: []byte(`{"country": "brazil"}`),
				rw:   newNoopWriter(false),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh := NewHTTPRequestHandler(tt.args.rw, nil)
			err := rh.Write(tt.args.data)
			if tt.wantErr {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}
