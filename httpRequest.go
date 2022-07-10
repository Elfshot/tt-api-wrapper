package ttRequests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var client = &http.Client{}

func Get_(url string, strucVar any) (*any, error) {
	res, resErr := request(url)
	if resErr != nil {
		return nil, resErr
	}

	_, parseErr := parse(res, strucVar)
	if parseErr != nil {
		return nil, parseErr
	}

	return &strucVar, nil
}

func GetNoParse_(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Tycoon-Key", os.Getenv("TYCOON_KEY_PRIVATE"))

	resp, err := client.Do(req)
	if err != nil { // Pass on any HTTP error
		fmt.Println(err)
		return nil, errors.New("HTTP Error: " + err.Error())
	}

	return resp, nil
}

func parse(body string, parsed any) (*any, error) {
	err := json.Unmarshal([]byte(body), &parsed) // Parse body string into the struct
	if err != nil {                              // Pass on any JSON error
		fmt.Println(err, body)
		return nil, errors.New("JSON Error: " + err.Error())
	}

	return &parsed, nil // Return reference to parsed struct | If "parsed" is not passed as reference, it will return an empty struct
}

func request(url string) (string, error) {
	// Make the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("X-Tycoon-Key", os.Getenv("TYCOON_KEY_PRIVATE"))

	resp, err := client.Do(req)
	if err != nil { // Pass on any HTTP error
		fmt.Println(err)
		return "", errors.New("HTTP Error: " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body) // Read the body into a string
	if err != nil {                        // Pass on any READ error
		fmt.Println(err)
		return "", errors.New("READ Error: " + err.Error())
	}
	bodyContent := string(body)

	if resp.StatusCode != 200 { // Pass on any HTTP error codes
		fmt.Printf("Error code: %d.\nContent: %s\n", resp.StatusCode, bodyContent)
		return "", errors.New("HTTP Error: " + resp.Status)
	}

	return bodyContent, nil
}
