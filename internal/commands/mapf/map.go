package mapf

import (
	"fmt"
	"log"

	"github.com/Woutjeee/go_pokedex/internal/client"
	"github.com/Woutjeee/go_pokedex/internal/commands/types"
)

type LocationAreasResponse struct {
	Results []types.Area `json:"results"`
}

func GetMap() error {
	var locationAreasResponse LocationAreasResponse
	err := client.Get("location-area", &locationAreasResponse, len(locationAreasResponse.Results))
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, location := range locationAreasResponse.Results {
		fmt.Println(location.Name)
	}

	fmt.Println("Length of results is", len(locationAreasResponse.Results))

	return nil
}
