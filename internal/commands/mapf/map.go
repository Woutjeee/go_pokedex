package mapf

import (
	"fmt"
	"log"

	"github.com/Woutjeee/go_pokedex/internal/client"
	"github.com/Woutjeee/go_pokedex/internal/commands/types"
)

type LocationAreasResponse struct {
	Results []types.Area `json:"results"`
	Next    string       `json:"next"`
}

var firstApiPath string = "https://pokeapi.co/api/v2/location-area"

func GetMap() error {
	var locationAreasResponse LocationAreasResponse

	err := client.Get(firstApiPath, &locationAreasResponse)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, location := range locationAreasResponse.Results {
		fmt.Println(location.Name)
	}

	firstApiPath = locationAreasResponse.Next
	return nil
}
