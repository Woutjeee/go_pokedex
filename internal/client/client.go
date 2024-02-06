package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Get[T any](path string, t *T) error {
	var resp *http.Response
	var err error

	resp, err = http.Get(path)

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
	// jsonResponse, err := json.MarshalIndent(t, "", "    ")
	// if err != nil {
	// 	return err
	// }
	// log.Printf("JSON Response:\n%s\n", jsonResponse)

	return nil
}
