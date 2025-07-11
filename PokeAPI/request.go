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
const BASELOCATIONAREA string = "https://pokeapi.co/api/v2/location-area/"

var mapLocationCache = pokecache.NewCache(10 * time.Minute)
var areaCache = pokecache.NewCache(24 * time.Hour)

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

	if res.StatusCode == 404 {
		return errors.New("request not found, possible area/pokemon does not exist")
	}

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

		if cacheData, hit := mapLocationCache.Get(locationsData.locationsEndpoint); hit {
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
		mapLocationCache.Add(locationsData.locationsEndpoint, data)

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

	if cacheData, hit := mapLocationCache.Get(*locationsData.prevLocationsEndpoint); hit {
		locCached := cacheData.(locations)
		fmt.Println("CACHED")

		if locCached.Next != nil {
			locationsData.locationsEndpoint = *locCached.Next
		} else {
			locationsData.locationsEndpoint = ""
		}
		locationsData.prevLocationsEndpoint = locCached.Prev
		return locCached, nil
	}

	var data locations
	if err := generalRequest(locationsData.locationsEndpoint, &data); err != nil {
		return locations{}, err
	}
	mapLocationCache.Add(*locationsData.prevLocationsEndpoint, data)

	if data.Next != nil {
		locationsData.locationsEndpoint = *data.Next
	} else {
		locationsData.locationsEndpoint = ""
	}
	locationsData.prevLocationsEndpoint = data.Prev
	return data, nil

}

func ExploreRequest(location string) (area, error) {
	fullURL := BASELOCATIONAREA + location
	if cacheData, hit := areaCache.Get(fullURL); hit {
		areaCache := cacheData.(area)
		fmt.Println("CACHED")
		return areaCache, nil
	}

	var data area
	if err := generalRequest(fullURL, &data); err != nil {
		return area{}, err
	}

	areaCache.Add(fullURL, data)
	return data, nil

}
