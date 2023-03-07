package handler

// Handler interface defines the following functions:
//
//	Read() ([]byte, error) ->
//		This function helps in reading from a specific source.
//		It returns the data read in form of bytes else an error describing what went wrong.
//
//	Write([]byte) error
//		This function helps in writing data.
//		It returns an error describing what went wrong.
type Handler interface {
	// Read function returns the bytes read or an error in case something went wrong.
	Read() ([]byte, error)

	// Write function writes the bytes of data and returns an error in case something went wrong.
	Write([]byte) error
}
