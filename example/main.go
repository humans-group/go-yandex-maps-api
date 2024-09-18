/*This main.go file is made for test/example purporse only.
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
	suggestAPI.AddLanguage(lang)
	suggestAPI.AddSearchPoint(lat, lng)
	suggestAPI.AddLimit(limit)
	suggestion, err := client.Suggest(suggestAPI, address)
	if err != nil {
		log.Fatalf("Suggest error: %v", err)
	}
	fmt.Printf("Suggest for text is %v\n", suggestion)
	fmt.Printf("Suggested len are: %d\n", len(suggestion.Results))
	fmt.Printf("Suggested 0 result are: %s\n", suggestion.Results[0].Title.Text)
}


