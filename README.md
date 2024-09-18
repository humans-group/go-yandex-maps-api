# go-yandex-maps-api
Golang SDK for Yandex Maps API

## Usage
To use package import:
```import (
	"fmt"
	"log"

	"github.com/humans-grpoup/go-yandex-maps-api/services/suggest"
	"github.com/humans-grpoup/go-yandex-maps-api/services/httpclient"
)```go

Make sure that you have an API key before utilizing the package
``` const (
	apiSuggest = "<API_KEY>"//Need to be substituted by your own API key
	...
) ```go

Then you can make a request to Yandex Suggest API as follows:

```package main

import (
	"fmt"
	"log"

	"github.com/humans-grpoup/go-yandex-maps-api/services/suggest"
	"github.com/humans-grpoup/go-yandex-maps-api/services/httpclient"
)

const (
	apiSuggest = "<API_KEY>"//Need to be substituted by your own API key
	lat = 69.02
	lng = 42.01
	lang = "eng"
	limit = 5
)

func main() {
	// Example: Suggest 
	address := "New York"
	suggestAPI := suggest.NewSuggest(apiSuggest)
    // Build a URL for the API request
    // Add a response language
	suggestAPI.AddLanguage(lang)
    // Add a search point coordinate
	suggestAPI.AddSearchPoint(lat, lng)
    // Add a limit on response results
	suggestAPI.AddLimit(limit)
    //
	suggestion, err := client.Suggest(suggestAPI, address)
	if err != nil {
		log.Fatalf("Suggest error: %v", err)
	}
	fmt.Printf("Suggest for text is %v\n", suggestion)
}```go

Sturcure in retrieving by Yandex Suggest is:
```	SuggestResponse struct {
		SuggestReqID string         `json:"suggest_reqid"`
		Results      []SuggestResult `json:"results"`
	}
	
	SuggestResult struct {
		Title    Title    `json:"title"`
		Subtitle Subtitle `json:"subtitle,omitempty"`
		Tags     []string `json:"tags"`
		Distance Distance `json:"distance"`
	}
	
	Title struct {
		Text string    `json:"text"`
		HL   []HLRange `json:"hl,omitempty"`
	}
	
	HLRange struct {
		Begin int `json:"begin"`
		End   int `json:"end"`
	}
	
	Subtitle struct {
		Text string `json:"text"`
	}
	
	Distance struct {
		Value float64 `json:"value"`
		Text  string  `json:"text"`
	}```go


## License
MIT
