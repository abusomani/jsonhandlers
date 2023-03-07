package handler

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpResponseRead(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		err     error
		wantRes bool
		res     []byte
		prepare func() *http.Response
	}{
		{
			name:    "should return error if the response is nil",
			wantErr: true,
			err:     errors.New("response is nil"),
			wantRes: false,
			res:     nil,
			prepare: func() *http.Response {
				return nil
			},
		},
		{
			name:    "should return error if the response body is nil",
			wantErr: true,
			err:     errors.New("response body is nil"),
			wantRes: false,
			res:     nil,
			prepare: func() *http.Response {
				return &http.Response{}
			},
		},
		{
			name:    "should return error if reading the response body fails",
			wantErr: true,
			err:     errors.New("error while reading body: test error"),
			wantRes: false,
			res:     nil,
			prepare: func() *http.Response {
				return &http.Response{
					Body: io.NopCloser(errReader(0)),
				}
			},
		},
		{
			name:    "should not return error even if closing the response body fails",
			wantErr: false,
			err:     nil,
			wantRes: true,
			res:     []byte(`{"country": "brazil"}`),
			prepare: func() *http.Response {
				return &http.Response{
					Body: newMockReader(bytes.NewReader([]byte(`{"country": "brazil"}`))),
				}
			},
		},
		{
			name:    "should return the bytes read from the response",
			wantErr: false,
			err:     nil,
			wantRes: true,
			res:     []byte(`{"country": "brazil"}`),
			prepare: func() *http.Response {
				return &http.Response{
					Body: ioutil.NopCloser(bytes.NewReader([]byte(`{"country": "brazil"}`))),
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var response *http.Response
			if tt.prepare != nil {
				response = tt.prepare()
			}
			rh := NewHTTPResponseHandler(response)
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

func TestHttpResponseWrite(t *testing.T) {
	type Args struct {
		data []byte
	}
	tests := []struct {
		name    string
		wantErr bool
		args    Args
		err     error
	}{
		{
			name:    "should return error if write method is invoked",
			wantErr: true,
			err:     errors.New("not allowed"),
			args: Args{
				data: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh := NewHTTPResponseHandler(nil)
			err := rh.Write(tt.args.data)
			if tt.wantErr {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}
