package main

import (
	"errors"
	"fmt"
	// "github.com/stretchr/testify/assert"
	// "testing"
)

func MustUnique(a []int, b []int) (int, error) {

    if len(a) == 0  || len(b) == 0 {
        fmt.Print("input is empty")
        return 0, errors.New("input is empty")
    }
    if len(a) >= 10  || len(b) >= 10 {
        fmt.Print("input too long")
		return 0, errors.New("input too long")
    }

	fmt.Print(len(a), len(b))
	return 1, errors.New("Nice")
}

type Person struct {
    Name string
    Surname string
    Age int
    Hobbies []string
}

func Validate(people []Person) []error {

}

// func Test_Example(t *testing.T) {
//     assert.Equal(t, true, true)
// }

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{2, 3, 4, 5}
	MustUnique(a, b)
}
