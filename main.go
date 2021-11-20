package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

const apiPath = "https://mach-eight.uc.r.appspot.com/"

type Measure struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Hin       float64 `json:"h_in,string"`
	HMetter   float64 `json:"h_metters,string"`
}

type ResponseMeasure struct {
	Values []Measure `json:"values"`
}

func main() {
	// get user input
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

	// make http request to get json
	response, err := getJSONValues()

	if err != nil {
		fmt.Println(err)
		return
	}

	// fill only numbers unrepeat
	numbers := make(map[float64][]int)

	for i := 0; i < len(response.Values); i++ {
		first := response.Values[i]
		numbers[first.Hin] = append(numbers[first.Hin], i)
	}

	keys := make([]float64, 0, len(numbers))
	for k := range numbers {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	min := keys[0]
	max := keys[len(keys)-1]
	max_2 := keys[len(keys)-2]

	if inputUser <= int(min+min) {
		fmt.Println("No matches found")
		return
	} else if inputUser > int(max + max_2) {
		fmt.Println("No matches found")
		return
	}

	// find pair that have matches inch sum with input user
	numbersValidated := make(map[float64]bool)
	isMatch := false
	counter := 0
	for key, _ := range numbers {
		for key2, _ := range numbers {
			if numbersValidated[key] || numbersValidated[key2] {
				continue
			}
			if key+key2 == float64(inputUser) {
				// fmt.Println(key, " ", key2)
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

func GET(url string) (*http.Response, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func getJSONValues() (*ResponseMeasure, error) {
	res, err := GET(apiPath)

	if res == nil {
		return nil, fmt.Errorf("Failed to connect with api")
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Failed to get json")
	}

	defer res.Body.Close()
	var response ResponseMeasure

	// convert bytes of body to responseDTO
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("Failed to decode json")
	}

	return &response, nil

}
