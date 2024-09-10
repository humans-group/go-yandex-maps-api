package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"github.com/humans-grpoup/go-yandex-maps-api/services/geocoder"
	"github.com/humans-grpoup/go-yandex-maps-api/services/suggest"
	"github.com/valyala/fasthttp"
)

const baseURL = "https://geocode-maps.yandex.ru/1.x/"
const apiKey = "1b203825-74f8-4936-ab55-4a5190fd7d09"
const apiSuggest = "289b3faf-6a1e-4f2b-b389-77f6bf9c7547"


func ForwardGeocode(address string) (string, error) {
	params := url.Values{}
	params.Set("apikey", apiKey)
	params.Set("geocode", address)
	params.Set("format", "json")

	url := baseURL + "?" + params.Encode()

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		return "", err
	}

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	var geocodeResponse geocoder.GeocodeResponse
	if err := json.Unmarshal(resp.Body(), &geocodeResponse); err != nil {
		return "", err
	}

	if len(geocodeResponse.Response.GeoObjectCollection.FeatureMember) == 0 {
		return "", fmt.Errorf("no result found for address: %s", address)
	}

	coordinates := geocodeResponse.Response.GeoObjectCollection.FeatureMember[0].GeoObject.Point.Pos
	return coordinates, nil
}

func ReverseGeocode(lat, lon string) (string, error) {
	coords := fmt.Sprintf("%s,%s", lon, lat)

	params := url.Values{}
	params.Set("apikey", apiKey)
	params.Set("geocode", coords)
	params.Set("format", "json")

	url := baseURL + "?" + params.Encode()

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		return "", err
	}

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	var geocodeResponse geocoder.GeocodeResponse
	if err := json.Unmarshal(resp.Body(), &geocodeResponse); err != nil {
		return "", err
	}

	if len(geocodeResponse.Response.GeoObjectCollection.FeatureMember) == 0 {
		return "", fmt.Errorf("no result found for coordinates: %s, %s", lat, lon)
	}

	address := geocodeResponse.Response.GeoObjectCollection.FeatureMember[0].GeoObject.Name
	return address, nil
}

func Suggest(apiKey, text string, suggestObj *suggest.Suggest) ([]string, error) {
	url := suggestObj.BaseURL(apiKey, text)
	fmt.Println("Suggest url is", url)
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	results := make([]string,0)
	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		return results, err
	}

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	var suggestResponse suggest.SuggestResponse
	if err := json.Unmarshal(resp.Body(), &suggestResponse); err != nil {
		return results, err
	}

	if len(suggestResponse.Results) == 0 {
		return results, fmt.Errorf("no result found for suggestion: %s", text)
	}
	for _, result := range suggestResponse.Results {
        results = append(results, result.Title.Text)
    }
	return results, nil

}

func main() {
	// Example: Forward Geocoding
	address := "New York"
	coordinates, err := ForwardGeocode(address)
	if err != nil {
		log.Fatalf("Forward geocoding error: %v", err)
	}
	fmt.Printf("Coordinates for '%s': %s\n", address, coordinates)

	// Example: Reverse Geocoding
	lat := "25.2535"
	lon := "33.6504"
	reverseAddress, err := ReverseGeocode(lat, lon)
	if err != nil {
		log.Fatalf("Reverse geocoding error: %v", err)
	}
	fmt.Printf("Address for coordinates (%s, %s): %s\n", lat, lon, reverseAddress)
	fmt.Println("Suggest for")
	suggestion, err := Suggest(apiSuggest, "Burj", suggest.NewSuggest("en_US", 10))
	if err != nil {
		log.Fatalf("Suggest error: %v", err)
	}
	fmt.Printf("Suggest for text is %v", suggestion)

}


