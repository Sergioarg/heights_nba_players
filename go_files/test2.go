package main

import (
	"fmt"
	"unicode"
)

type Person struct {
	Name    string
	Surname string
	Age     int
	Hobbies []string
}

func IsUpper(s string) bool {
	/* Check if string is uppercase */
    for _, r := range s {
        if !unicode.IsUpper(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
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

	if p.Name == "" || p.Surname == "" {
		errors += "\nName or Surname cannot be empty"
	} else if IsUpper(string(p.Name[0])) == false || IsUpper(string(p.Surname[0])) == false {
		errors += "\nName or Surname does no start with capital letter"
	}
	if p.Age <= 0 {
		errors += "\nAge is lower than 0"
	} else if p.Age >= 120 {
		errors += "\nAge is higher than 120"
	}

	for i := 1; i < len(p.Hobbies); i++ {
		if p.Hobbies[i] == p.Hobbies[0] {
			errors += "Hobbies are not unique"
		}
	}

	return errors
}


func main() {
	people := []Person {
		{Name: "William", Age: 29, Surname: "Andres", Hobbies: []string{"dance", "music"}},
		{Name: "Sergio", Age: 21, Surname: "Ramos", Hobbies: []string{"dance", "music", "sleep"}},
	}
	errors := validate(people)
	fmt.Println(errors)
}
