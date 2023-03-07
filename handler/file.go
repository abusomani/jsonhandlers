package handler

type FileHandler struct {
}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

func (fh *FileHandler) Parse(data []byte, v any) {

}
