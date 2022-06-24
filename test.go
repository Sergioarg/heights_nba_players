package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

func MostUnique(a []int, b []int) (int, error) {
	const error_msg = "input is empty"
	const error_long = "input is so long"
	if len(a) == 0 || len(b) == 0 {
		return 0, errors.New(error_msg)
	}

	if len(a) > 10 || len(b) == 10 {
		return 0, errors.New(error_long)
	}

	aMax := float64(uniqueMax(a))
	bMax := float64(uniqueMax(b))

	return int(math.Max(aMax, bMax)), nil
}

func uniqueMax(a []int) int {
	m1 := make(map[int]int)

	for _, elem := range a {
		m1[elem]++
	}

	var uniq1 []int

	for k, v := range m1 {
		if v == 1 {
			uniq1 = append(uniq1, k)
		}
	}

	if len(uniq1) == 0 {
		return 0
	}

	sort.Ints(uniq1)
	return uniq1[len(uniq1)-1]
}

type Person struct {
	Name    string
	Surname string
	Age     int
	Hobbies []string
}

func main() {
	a := []int{3, 8, 3, 5, 5}
	b := []int{2, 2, 3, 3, 2, 4}

	fmt.Println(MostUnique(a, b))
}
