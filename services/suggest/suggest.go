package suggest

import (
	"fmt"
)

// SuggestResponse represents the full response from the Yandex Suggest API.

const (
	// YandexSuggestAPI is the base URL for the Yandex Suggest API.
	yandexSuggestAPI = "http://suggest-maps.yandex.ru/v1/suggest"
)

type (
	SuggestAPI struct {
		endpoint   string
		parameters string
	}
)

func NewSuggestAPI(apiKey string, baseURLs ...string) SuggestAPI {
	return SuggestAPI{
		endpoint:   getURL(apiKey, baseURLs...),
		parameters: "",
	}
}

func getURL(apiKey string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return fmt.Sprintf("%s?format=json&apikey=%s", yandexSuggestAPI, apiKey)
}

func (s SuggestAPI) AddSearchPoint(lat, lng float64) SuggestAPI {
	return SuggestAPI{endpoint: s.endpoint,
		parameters: fmt.Sprintf("%s&ll=%.6f,%.6f", s.parameters, lat, lng)}
}

func (s SuggestAPI) AddLanguage(lang string) SuggestAPI {
	return SuggestAPI{endpoint: s.endpoint,
		parameters: fmt.Sprintf("%s&lang=%s", s.parameters, lang)}
}

func (s SuggestAPI) AddLimit(limit int) SuggestAPI {
	return SuggestAPI{endpoint: s.endpoint,
		parameters: fmt.Sprintf("%s&results=%d", s.parameters, limit)}
}

func (s SuggestAPI) AddBoundaryBox(minlng, minlat, maxlng, maxlat float64) SuggestAPI {
	return SuggestAPI{endpoint: s.endpoint,
		parameters: fmt.Sprintf("%s&bbox=%.6f,%.6f~%.6f,%.6f", s.parameters, minlng, minlat, maxlng, maxlng)}
}
func (s *SuggestAPI) GeosuggestURL(address string) string {
	URL := string(s.endpoint) + s.parameters + "&text=" + address
	return URL
}
