package geocode

import (
	"fmt"
	"net/url"
)

const (
	// YandexGeocodeAPI is the base URL for the Yandex Geocode API.
	yandexGeocodeAPI = "https://geocode-maps.yandex.ru/1.x"
)

type (
	baseURL    string
	GeocodeAPI struct {
		endpoint   baseURL
		parameters string
	}
)

func NewGeocodeAPI(apiKey string, baseURLs ...string) GeocodeAPI {
	return GeocodeAPI{
		endpoint:   baseURL(getURL(apiKey, baseURLs...)),
		parameters: "",
	}
}

func getURL(apiKey string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return fmt.Sprintf("%s?format=json&apikey=%s", yandexGeocodeAPI, apiKey)
}

func (s GeocodeAPI) AddSearchPoint(lat, lng float64) GeocodeAPI {
	return GeocodeAPI{endpoint: s.endpoint,
		parameters: fmt.Sprintf("%s&ll=%.6f,%.6f", s.parameters, lat, lng)}
}

func (s GeocodeAPI) AddLanguage(lang string) GeocodeAPI {
	return GeocodeAPI{endpoint: s.endpoint,
		parameters: fmt.Sprintf("%s&lang=%s", s.parameters, lang)}
}

func (s GeocodeAPI) AddLimit(limit int) GeocodeAPI {
	return GeocodeAPI{endpoint: s.endpoint,
		parameters: fmt.Sprintf("%s&results=%d", s.parameters, limit)}
}

func (s GeocodeAPI) AddBoundaryBox(minlng, minlat, maxlng, maxlat float64) GeocodeAPI {
	return GeocodeAPI{endpoint: s.endpoint,
		parameters: fmt.Sprintf("%s&bbox=%.6f,%.6f~%.6f,%.6f", s.parameters, minlng, minlat, maxlng, maxlng)}
}

func (s *GeocodeAPI) ForwardGeocodeURL(address string) string {
	URL := string(s.endpoint) + s.parameters + "&geocode=" + address
	return URL
}

func (s *GeocodeAPI) ReverseGeocodeURL(lat, lng float64) string {
	URL := string(s.endpoint) + s.parameters + "&sco=latlong&geocode=" + url.QueryEscape(fmt.Sprintf("%f,%f", lat, lng))
	return URL
}
