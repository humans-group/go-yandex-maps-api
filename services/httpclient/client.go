package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/humans-group/go-yandex-maps-api/services/geocode"
	"github.com/humans-group/go-yandex-maps-api/services/suggest"
)

// DefaultTimeout for the request execution
const DefaultTimeout = time.Second * 8

// ErrTimeout occurs when no response returned within timeoutInSeconds
var ErrTimeout = errors.New("TIMEOUT")

type (
	// EndpointBuilder defines functions that build urls for geosuggest
	EndpointBuilder interface {
		GeosuggestURL(address string) string
		AddSearchPoint(lat, lng float64)
		AddLanguage(lang string)
		AddLimit(limit int)
	}

	HTTPClient interface {
		Execute(ctx context.Context, url string, obj interface{}) error
		GetTimeout() time.Duration
	}

	SimpleHTTPClient struct {
		Timeout time.Duration
	}
)

func Suggest(client HTTPClient, suggestAPI *suggest.SuggestAPI, text string) (*SuggestResponse, error) {
	return GeoRequest[SuggestResponse](client, suggestAPI.GeosuggestURL(text))
}

func ForwardGeocode(client HTTPClient, geocodeAPI *geocode.GeocodeAPI, text string) (*GeocodeResponse, error) {
	return GeoRequest[GeocodeResponse](client, geocodeAPI.ForwardGeocodeURL(text))
}

func ReverseGeocode(client HTTPClient, geocodeAPI *geocode.GeocodeAPI, lat, lng float64) (*GeocodeResponse, error) {
	return GeoRequest[GeocodeResponse](client, geocodeAPI.ReverseGeocodeURL(lat, lng))
}

// Generic function to handle different response types
func GeoRequest[T any](client HTTPClient, text string) (*T, error) {
	// Use the generic type T for the response
	responseParser := new(T)

	ctx, cancel := context.WithTimeout(context.TODO(), client.GetTimeout())
	defer cancel()

	// A custom struct to handle both response and error
	chResp := make(chan *T, 1)
	chErr := make(chan error, 1)
	// Goroutine to execute the client request
	go func(chResp chan *T, chErr chan error) {
		fmt.Println("text is ", text, url.QueryEscape(text))
		err := client.Execute(ctx, url.QueryEscape(text), responseParser)
		if err != nil {
			chErr <- err
		} else {
			chResp <- responseParser
		}
	}(chResp, chErr)

	// Handle the result or timeout
	select {
	case <-ctx.Done():
		return nil, ErrTimeout
	case res := <-chResp:
		return res, nil
	case errResponse := <-chErr:
		return nil, errResponse
	}
}

func (sh SimpleHTTPClient) Execute(ctx context.Context, url string, obj interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		ErrLogger.Printf("cannot implement request %e: \n", err)
		return err
	}

	body := strings.Trim(string(data), " []")
	if resp.StatusCode != http.StatusOK {
		ErrLogger.Printf("Received status code %d: %s\n", resp.StatusCode, body)
		return fmt.Errorf("received status code %d", resp.StatusCode)
	}
	DebugLogger.Printf("Received response: %s\n", body)
	if body == "" {
		return nil
	}
	if err := json.Unmarshal([]byte(body), obj); err != nil {
		ErrLogger.Printf("Error unmarshalling response: %s\n", err.Error())
		return err
	}

	return nil
}

func (sh SimpleHTTPClient) GetTimeout() time.Duration {
	return DefaultTimeout
}
