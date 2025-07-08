package pokeapi

import (
	"encoding/json"
	"errors"
	"net/http"
)

const BASEURL string = "https://pokeapi.co/api/v2/location-area/"

var ErrorBackPage = errors.New("you are on the backpage")
var ErrorPageFront = errors.New("you are on the front page")

var LocationsData = struct {
	locationsEndpoint     string
	PrevLocationsEndpoint *string
}{
	locationsEndpoint:     BASEURL,
	PrevLocationsEndpoint: nil,
}

func generalRequest[T any](url string, buffer *T) error {

	res, err := http.Get(LocationsData.locationsEndpoint)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(buffer); err != nil {
		return err
	}

	return nil

}

func MapRequest() (locations, error) {
	if LocationsData.locationsEndpoint != "" {
		var data locations
		if err := generalRequest(LocationsData.locationsEndpoint, &data); err != nil {
			return locations{}, err
		}

		*(LocationsData.PrevLocationsEndpoint) = LocationsData.locationsEndpoint

		if data.Next == nil {
			LocationsData.locationsEndpoint = ""
		} else {
			LocationsData.locationsEndpoint = *data.Next
		}

		return data, nil

	}

	return locations{}, ErrorBackPage

}
