package handler

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name     string
		wantErr  bool
		err      error
		wantRes  bool
		res      []byte
		fileName string
		prepare  func()
		cleanup  func()
	}{
		{
			name:     "should return error if the file does not exist",
			wantErr:  true,
			err:      errors.New(" error in opening file : open dummy.json: no such file or directory"),
			wantRes:  false,
			res:      nil,
			fileName: "dummy.json",
		},
		{
			name:     "should return the bytes read from the file",
			wantErr:  false,
			err:      nil,
			wantRes:  true,
			res:      []byte(`{"country": "brazil"}`),
			fileName: "test.json",
			prepare: func() {
				os.WriteFile("test.json", []byte(`{"country": "brazil"}`), os.ModePerm)
			},
			cleanup: func() {
				os.Remove("test.json")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepare != nil {
				tt.prepare()
			}
			f := NewFileHandler(tt.fileName)
			res, err := f.Read()
			if tt.wantErr {
				assert.Equal(t, tt.err, err)
			}
			if tt.wantRes {
				assert.Equal(t, tt.res, res)
			}
			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

func TestWrite(t *testing.T) {
	tests := []struct {
		name     string
		arg      []byte
		wantErr  bool
		err      error
		fileName string
		prepare  func()
		cleanup  func()
	}{
		{
			name:     "should not return any error upon successfully writing",
			arg:      []byte(`{"country": "brazil"}`),
			wantErr:  false,
			err:      nil,
			fileName: "test.json",
			prepare: func() {
				os.Create("test.json")
			},
			cleanup: func() {
				os.Remove("test.json")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepare != nil {
				tt.prepare()
			}
			f := NewFileHandler(tt.fileName)
			err := f.Write(tt.arg)
			if tt.wantErr {
				assert.Equal(t, tt.err, err)
			}
			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}
