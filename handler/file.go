package handler

import (
	"fmt"
	"os"
)

// FileHandler implements the Handler interface
// It takes in a fileName which specifies the file on which the operations have to be performed.
type FileHandler struct {
	fileName string
}

// NewFileHandler takes in a fileName
// It returns a FileHandler instance.
func NewFileHandler(fileName string) *FileHandler {
	return &FileHandler{fileName: fileName}
}

// Read function returns the bytes read from the given file or an error in case something went wrong.
func (fh *FileHandler) Read() ([]byte, error) {
	data, err := os.ReadFile(fh.fileName)
	if err != nil {
		return nil, fmt.Errorf(" error in opening file : %s", err.Error())
	}
	return data, nil
}

// Write function writes the bytes of data to the given file and returns an error in case something went wrong.
func (fh *FileHandler) Write(input []byte) error {
	return os.WriteFile(fh.fileName, input, os.ModePerm)
}
