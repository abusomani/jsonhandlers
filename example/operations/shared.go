package operations

import "log"

const (
	testFilePath = "example/operations/test.json"
)

type student struct {
	ID     int
	Name   string
	Branch string
}

type school struct {
	Students []student
}

func handleError(text string, err error) {
	if err != nil {
		log.Fatalf(text, err.Error())
	}
}
