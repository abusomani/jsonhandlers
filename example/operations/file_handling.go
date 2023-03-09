package operations

import (
	"fmt"

	"github.com/abusomani/jsonhandlers"
)

func GetStudentsFromFile() []student {
	return handleFile()
}

func handleFile() []student {
	jh := jsonhandlers.New(jsonhandlers.WithFileHandler(testFilePath))

	var sch school
	err := jh.Unmarshal(&sch)
	handleError("error in unmarshalling %s", err)
	fmt.Printf("School info is : %+v\n", sch)

	// add a new student to the school
	sch.Students = append(sch.Students[:2], student{
		ID:     3,
		Name:   "The new student",
		Branch: "AI",
	})

	err = jh.Marshal(sch)
	handleError("error in marshalling %s", err)
	fmt.Printf("Updated school info after admission of new student is : %+v\n", sch)

	// remove the new student as he was very mischievous
	sch.Students = sch.Students[:2]

	err = jh.Marshal(sch)
	handleError("error in marshalling %s", err)
	fmt.Printf("Updated school info after retaining all good students is : %+v\n", sch)
	return sch.Students
}
