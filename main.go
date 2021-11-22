package main

import "fmt"

const apiPath = "https://mach-eight.uc.r.appspot.com/"

/**
 * struct Measure - struct for order the nba players
 * @FirstName: firs name
 * @LastName: last name
 * @Hin: height in inches
 * @HMetter: height in meters
 */

type Measure struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Hin       float64 `json:"h_in,string"`
	HMetter   float64 `json:"h_metters,string"`
}

/**
 * struct ResponseMeasure - struct for order the nba players
 * @Measure: Array for storage the results of json values
 */
type ResponseMeasure struct {
	Values []Measure `json:"values"`
}

/**
 * getInputUser - check and get the input of the user
 * @response: structure that has the value
 *
 * Return: the number input of the user
 */
func getInputUser() int {

	var inputUser int

	for {

		fmt.Print("Introduce Inch: ")
		fmt.Scanf("%d", &inputUser)

		if inputUser == 0 {
			fmt.Println("Introduce valid input")
		} else {
			break
		}
	}
	return inputUser
}


// main - Entry point
func main() {

	input := getInputUser()
	response, err := getJSONValues()

	if err != nil {
		fmt.Println(err)
		return
	}

	numbers := getUniqueInch(response)
	checkLimitsInch(numbers, input)
	findPairsWhitMatchSumInch(numbers, response, input)
}
