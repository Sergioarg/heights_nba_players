package main

import (
	"fmt"
	"testing"
	"os"
)

// TestApiInch - test base case
func TestApiInch(t *testing.T) {

	input := 139
	fmt.Println("Input test:", input)
	response, err := getJSONValues()

	if err != nil {
		fmt.Println(err)
		return
	}

	numbers := getUniqueInch(response)
	checkLimitsInch(numbers, input)
	findPairsWhitMatchSumInch(numbers, response, input)
	fmt.Println("")
}

// TestApiMinInch - test with the sum of the two higher values
func TestApiMinInch(t *testing.T) {

	input := 1
	fmt.Println("Input test:", input)
	response, err := getJSONValues()

	if err != nil {
		fmt.Println(err)
		return
	}

	numbers := getUniqueInch(response)
	if os.Getenv("EXIT_STATUS") == "1" {
		checkLimitsInch(numbers, input)
        return
    }
	findPairsWhitMatchSumInch(numbers, response, input)
	fmt.Println("")
}

// TestApiMaxInch - test with the max values in inches
func TestApiMaxInch(t *testing.T) {

	input := 200
	fmt.Println("Input test:", input)
	response, err := getJSONValues()

	if err != nil {
		fmt.Println(err)
		return
	}

	numbers := getUniqueInch(response)
	if os.Getenv("EXIT_STATUS") == "1" {
		checkLimitsInch(numbers, input)
        return
    }
	findPairsWhitMatchSumInch(numbers, response, input)
	fmt.Println("")
}

// TestApiNegativeInch - test with negative values in inches
func TestApiNegativeInch(t *testing.T) {

	input := -50
	fmt.Println("Input test:", input)
	response, err := getJSONValues()

	if err != nil {
		fmt.Println(err)
		return
	}

	numbers := getUniqueInch(response)
	if os.Getenv("EXIT_STATUS") == "1" {
		checkLimitsInch(numbers, input)
        return
    }
	findPairsWhitMatchSumInch(numbers, response, input)
}
