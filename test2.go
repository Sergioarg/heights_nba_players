package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Surname string
	Age     int
	Hobbies []string
}

func validate(people []Person) []error {
	var listErrors []error

	for index, p := range people {
		if msg := validateOnly(p); msg != "" {
			err := fmt.Errorf("Person: %d, %s", index, msg)
			listErrors = append(listErrors, err)
		}
	}
	return listErrors
}

func validateOnly(p Person) string {
	errors := ""
	if p.Name == "" {
		errors += ""
	}

	return errors
}

func main() {
	people := []Person{{Name: "william", Age: 29, Surname: "Andres", Hobbies: []string{"dance", "music"}},
		{Name: "", Age: -29, Surname: "", Hobbies: []string{"dance", "music", "dance"}}}

	errors := validate(people)
	fmt.Println(errors)
}
