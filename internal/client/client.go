package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var base_url string = "https://pokeapi.co/api/v2/"

func Get[T any](path string, t *T) error {
	resp, err := http.Get(base_url + path)
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
	return nil
}
