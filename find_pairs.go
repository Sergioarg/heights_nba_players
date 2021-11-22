package main

import "fmt"


/**
 * findPairsWhitMatchSumInch - find pairs of the inch values
 * @numbers: maps with the values of th inch
 * @response: structure that has the value
 * @input: input insert for the user
 *
 * Return: nothing
 */
 func findPairsWhitMatchSumInch(numbers map[float64][]int, response *ResponseMeasure, input int) {

	numbersValidated := make(map[float64]bool)
	isMatch := false
	counter := 0

	for key, _ := range numbers {
		for key2, _ := range numbers {
			if numbersValidated[key] || numbersValidated[key2] {
				continue
			}
			if key+key2 == float64(input) {
				isMatch = true
				for _, v := range numbers[key] {
					value := response.Values[v]
					name1 := value.FirstName + " " + value.LastName

					for _, v := range numbers[key2] {
						value := response.Values[v]
						name2 := value.FirstName + " " + value.LastName

						if name1 != name2 {
							fmt.Printf("%s\t%s\n", name1, name2)
						}
					}
				}
				numbersValidated[key] = true
				numbersValidated[key2] = true
			}
			counter++
		}
		numbersValidated[key] = true
	}

	if !isMatch {
		fmt.Println("No matches found")
	}
}
