package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var baseURL string = "https://pokeapi.co/api/v2/"

func Get[T any](path string, t *T, skip ...int) error {
	var resp *http.Response
	var err error

	if len(skip) > 0 && skip[0] != 0 {
		resp, err = http.Get(baseURL + path + "/" + strconv.Itoa(skip[0]))
	} else {
		resp, err = http.Get(baseURL + path)
	}

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(t)
	if err != nil {
		return err
	}

	// Print the JSON response to the console
	jsonResponse, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		return err
	}
	log.Printf("JSON Response:\n%s\n", jsonResponse)

	return nil
}
