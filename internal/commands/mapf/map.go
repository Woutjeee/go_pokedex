package mapf

import (
	"log"

	"github.com/Woutjeee/go_pokedex/internal/client"
	"github.com/Woutjeee/go_pokedex/internal/commands/types"
)

type LocationsResponse struct {
	Results []types.Location `json:"results"`
}

func GetMap() error {
	var locationsResponse LocationsResponse
	err := client.Get("location", &locationsResponse)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, location := range locationsResponse.Results {
		log.Printf("Location Name: %s", location.Name)
	}

	return nil
}
