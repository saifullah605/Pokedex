package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	pokecache "github.com/saifullah605/Pokedex/pokecache"
)

const BASELOCATIONS string = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

var locationCache = pokecache.NewCache(10 * time.Minute)

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

		if cacheData, hit := locationCache.Get(locationsData.locationsEndpoint); hit {
			locCached := cacheData.(locations)
			fmt.Println("CACHED")
			locationsData.locationsEndpoint = *locCached.Next
			locationsData.prevLocationsEndpoint = locCached.Prev
			return locCached, nil
		}

		var data locations
		if err := generalRequest(locationsData.locationsEndpoint, &data); err != nil {
			return locations{}, err
		}

		locationsData.prevLocationsEndpoint = data.Prev
		locationCache.Add(locationsData.locationsEndpoint, data)

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

	if cacheData, hit := locationCache.Get(*locationsData.prevLocationsEndpoint); hit {
		locCached := cacheData.(locations)
		fmt.Println("CACHED")
		// Set locationsEndpoint to the NEXT page, not the current one
		if locCached.Next != nil {
			locationsData.locationsEndpoint = *locCached.Next
		} else {
			locationsData.locationsEndpoint = ""
		}
		locationsData.prevLocationsEndpoint = locCached.Prev
		return locCached, nil
	}

	var data locations
	locationCache.Add(*locationsData.prevLocationsEndpoint, data)
	// Set locationsEndpoint to the NEXT page, not the current one
	if data.Next != nil {
		locationsData.locationsEndpoint = *data.Next
	} else {
		locationsData.locationsEndpoint = ""
	}
	locationsData.prevLocationsEndpoint = data.Prev
	return data, nil

}
