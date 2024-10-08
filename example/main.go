/*
This main.go file is made for test/example purporse only.
It sould be called with:
```

	go run ./example/main.go

```
with API key before
*/
package main

import (
	"fmt"
	"log"

	"github.com/humans-group/go-yandex-maps-api/services/geocode"
	client "github.com/humans-group/go-yandex-maps-api/services/httpclient"
	"github.com/humans-group/go-yandex-maps-api/services/suggest"
)

const (
	apiSuggest  = "<YOUR_SUGGEST_API_KEY>" //Need to be substituted by your own API key
	apiGeocoder = "<YOUR_GEOCODE_API_KEY>" //Need to be substituted by your own API key
	lat         = 69.02
	lng         = 42.01
	lang        = "en_US"
	limit       = 5
)

func main() {
	// Example: Suggest
	address := "New York"
	clientAPI := &client.SimpleHTTPClient{}
	suggestAPI := suggest.NewSuggestAPI(apiSuggest)
	suggestAPI = suggestAPI.
		AddSearchPoint(lat, lng).
		AddLimit(limit).
		AddLanguage(lang)
	suggestion, err := client.Suggest(clientAPI, &suggestAPI, address)
	if err != nil {
		log.Fatalf("Suggest error: %v", err)
	}
	fmt.Printf("Suggest for text is %v\n", suggestion)
	fmt.Printf("Suggested len are: %d\n", len(suggestion.Results))
	fmt.Printf("Suggested 0 result are: %s\n", suggestion.Results[0].Title.Text)
	geocodeAPI := geocode.NewGeocodeAPI(apiGeocoder)
	geocodeAPI = geocodeAPI.AddLanguage(lang).
		AddLimit(limit)
	forwardGeocodeResult, err := client.ForwardGeocode(clientAPI, &geocodeAPI, address)
	if err != nil {
		log.Fatalf("Forward geocode error: %v", err)
	}
	fmt.Printf("Forward geocode result for text is %v\n", forwardGeocodeResult)
	reverseGeocodeResult, err := client.ReverseGeocode(clientAPI, &geocodeAPI, lat, lng)
	if err != nil {
		log.Fatalf("Reverse geocode error: %v", err)
	}
	fmt.Printf("Reverse geocode result for text is %v\n", reverseGeocodeResult)
}
