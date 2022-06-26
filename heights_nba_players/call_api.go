package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
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
