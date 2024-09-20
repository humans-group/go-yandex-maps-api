package client

import (
    "context"
    "net/http"
	"net/url"
	"errors"
	"strings"
	"io"
	"time"
	"fmt"
	"encoding/json"

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
		EndpointBuilder
	}

	SimpleHTTPClient struct {
		Timeout time.Duration
		EndpointBuilder
	}
)

func Suggest(client HTTPClient, text string) (*SuggestResponse, error) {
	responseParser := &SuggestResponse{}

	ctx, cancel := context.WithTimeout(context.TODO(), client.GetTimeout())
	defer cancel()

	type sugResp struct {
		r *SuggestResponse
		e error
	}
	ch := make(chan sugResp, 1)

	go func(ch chan sugResp) {
		err := client.Execute(ctx, client.GeosuggestURL(url.QueryEscape(text)), responseParser)
		if err != nil {
			ch <- sugResp{
				r: nil,
				e: err,
			}
		}
		ch <- sugResp{
			r: responseParser,
			e: err,
		}
	}(ch)

	select {
		case <-ctx.Done():
			return nil, ErrTimeout
		case res := <-ch:
			return res.r, res.e
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
	if resp.StatusCode!= http.StatusOK {
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