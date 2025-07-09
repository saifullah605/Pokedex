package pokeapi

import (
	"encoding/json"
	"errors"
	"net/http"
)

const BASELOCATIONS string = "https://pokeapi.co/api/v2/location-area/"

var ErrorBackPage = errors.New("you are on the backpage")
var ErrorPageFront = errors.New("you are on the front page")

var locationsData = struct {
	locationsEndpoint     string
	prevLocationsEndpoint *string
}{
	locationsEndpoint:     BASELOCATIONS,
	prevLocationsEndpoint: nil,
}

func generalRequest[T any](url string, buffer *T) error {

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("bad status code: " + res.Status)
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(buffer); err != nil {
		return err
	}

	return nil

}

func MapRequest() (locations, error) {
	if locationsData.locationsEndpoint != "" {
		var data locations
		if err := generalRequest(locationsData.locationsEndpoint, &data); err != nil {
			return locations{}, err
		}

		locationsData.prevLocationsEndpoint = data.Prev

		if data.Next == nil {
			locationsData.locationsEndpoint = ""
		} else {
			locationsData.locationsEndpoint = *data.Next
		}

		return data, nil

	}

	return locations{}, ErrorBackPage

}

func PrevMapRequest() (locations, error) {
	if locationsData.prevLocationsEndpoint == nil {
		return locations{}, ErrorPageFront
	}

	var data locations
	if err := generalRequest(*locationsData.prevLocationsEndpoint, &data); err != nil {
		return locations{}, err
	}
	locationsData.locationsEndpoint = *locationsData.prevLocationsEndpoint
	locationsData.prevLocationsEndpoint = data.Prev
	return data, nil

}
