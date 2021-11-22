package main

import (
	"fmt"
	"os"
	"sort"
)
/**
 * getUniqueInch - get the unique values of the inch
 * @response: structure that has the value
 *
 * Return: numbers map[float64][]int
 */
 func getUniqueInch(response *ResponseMeasure) map[float64][]int {

	numbers := make(map[float64][]int)

	for i := 0; i < len(response.Values); i++ {
		first := response.Values[i]
		numbers[first.Hin] = append(numbers[first.Hin], i)
	}
	return numbers
}

/**
 * checkLimitsInch - check the max and min values of the inchs
 * @numbers: maps with the values of th inch
 * @input: input insert for the user
 *
 * Return: nothing
 */
func checkLimitsInch(numbers map[float64][]int, input int) {

	keys := make([]float64, 0, len(numbers))
	for k := range numbers {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	min := keys[0]
	max := keys[len(keys)-1]
	max_2 := keys[len(keys)-2]
	error_message := "No matches found"

	if input <= int(min+min) {
		fmt.Println(error_message)
		os.Exit(1)
	} else if input > int(max+max_2) {
		fmt.Println(error_message)
		os.Exit(1)
	}
}
