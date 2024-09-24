package suggest 
import (
	"fmt"
	"github.com/humans-group/go-yandex-maps-api/services/httpclient"
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
		parameters string
	}
)
	
func NewSuggest(apiKey string, baseURLs ...string) client.HTTPClient {
	return &client.SimpleHTTPClient{
		EndpointBuilder: &SuggestAPI {
			endpoint:  baseURL(getURL(apiKey, baseURLs...)),
			parameters: "",
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
	s.parameters = fmt.Sprintf("%s&ll=%.6f,%.6f", s.parameters, lat, lng)
}

func (s *SuggestAPI) AddLanguage(lang string) {
	s.parameters = fmt.Sprintf("%s&lang=%s", s.parameters, lang)
}

func (s *SuggestAPI) AddLimit(limit int) {
	s.parameters = fmt.Sprintf("%s&results=%d", s.parameters ,limit)
}

func (s *SuggestAPI) GeosuggestURL(address string) string {
	URL := string(s.endpoint) + s.parameters + "&text=" + address
	s.parameters = ""
	return URL
}
