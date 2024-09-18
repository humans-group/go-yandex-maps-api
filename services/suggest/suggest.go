package suggest 
import (
	"fmt"
	"github.com/humans-grpoup/go-yandex-maps-api/services/httpclient"
)
// SuggestResponse represents the full response from the Yandex Suggest API.

const (
    // YandexSuggestAPI is the base URL for the Yandex Suggest API.
    yandexSuggestAPI = "https://suggest-maps.yandex.ru/v1/suggest"
)
type (
	baseURL string
	SuggestAPI struct {
		endpoint baseURL 
	}
)
	
func NewSuggest(apiKey string, baseURLs ...string) client.HTTPClient {
	return &client.FastHTTPClient{
		EndpointBuilder: &SuggestAPI {
			endpoint:  baseURL(getURL(apiKey, baseURLs...)),
		},
	}
}

func getURL(apiKey string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return fmt.Sprintf("%s?format=json&apikey=%s", yandexSuggestAPI, apiKey)
}

func (s *SuggestAPI) AddSearchPoint(lat, lng float64) {
	s.endpoint = baseURL(fmt.Sprintf("%s&ll=%.6f,%.6f", string(s.endpoint), lat, lng))
}

func (s *SuggestAPI) AddLanguage(lang string) {
	s.endpoint = baseURL(fmt.Sprintf("%s&lang=%s",string(s.endpoint),lang))
}

func (s *SuggestAPI) AddLimit(limit int) {
	s.endpoint = baseURL(fmt.Sprintf("%s&results=%d",string(s.endpoint),limit))
}

func (s *SuggestAPI) GeosuggestURL(address string) string {
	return string(s.endpoint) + "&text=" + address
}
