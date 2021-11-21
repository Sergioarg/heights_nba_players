package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
)

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

/**
 * findPairsWhitMatchSumInch - find pairs of the inch values
 * @numbers: maps with the values of th inch
 * @response: structure that has the value
 * @input: input insert for the user
 *
 * Return: the response of the call
 */
func GET(url string) (*http.Response, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	return res, nil
}

/**
 * getJSONValues - get the json values calls from the API
 * @numbers: maps with the values of th inch
 * @response: structure that has the value
 * @input: input insert for the user
 *
 * Return: the response
 */
func getJSONValues() (*ResponseMeasure, error) {

	res, err := GET(apiPath)

	if res == nil {
		return nil, fmt.Errorf("failed to connect with api")
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get json")
	}

	defer res.Body.Close()
	var response ResponseMeasure

	// convert bytes of body to responseDTO
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode json")
	}

	return &response, nil

}

/**
 * main - Entry point
 */
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
