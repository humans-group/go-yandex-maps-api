package main_test

import (

	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/humans-group/go-yandex-maps-api/services/suggest"
	"github.com/humans-group/go-yandex-maps-api/services/httpclient"
	"github.com/stretchr/testify/assert"
)

var tokenSuggest = os.Getenv("YANDEX_SUGGEST_API_KEY")


func testServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(response))
	}))
}

func TestYandexGeoSuggest(t *testing.T) {
	ts := testServer(responseSugFound)
	defer ts.Close()
	suggestAPI := suggest.NewSuggest(tokenSuggest, ts.URL+"/")
	suggestAPI.AddLanguage("eng")
	suggestion, err := client.Suggest(suggestAPI, "Burj")
	
	assert.NoError(t, err)
	assert.True(t, len(suggestion.Results) > 0)
	assert.True(t, suggestion.Results[0].Title.Text=="Burji")
	assert.True(t, suggestion.Results[1].Title.Text=="Downtown Dubai")
	assert.True(t, suggestion.Results[2].Title.Text=="City of Borjomi")
}

func TestYandexGeoSuggestNoResult(t *testing.T) {
	ts := testServer(responseSugNotFound)
	defer ts.Close()
	suggestAPI := suggest.NewSuggest(tokenSuggest, ts.URL+"/")
	suggestAPI.AddLanguage("eng")
	suggestion, err := client.Suggest(suggestAPI, "$@#@")
	assert.Nil(t, err)
	assert.Nil(t, suggestion.Results)
}

const ( 
	responseSugFound = `{
  "suggest_reqid": "1726095099605352-4131673751-suggest-maps-yp-2",
  "results": [
    {
      "title": {
        "text": "Burji",
        "hl": [
          {
            "begin": 0,
            "end": 4
          }
        ]
      },
      "subtitle": {
        "text": "Kano"
      },
      "tags": [
        "locality"
      ],
      "distance": {
        "value": 1609715.142,
        "text": "998.02 mi"
      }
    },
    {
      "title": {
        "text": "Downtown Dubai"
      },
      "subtitle": {
        "text": "Dubai"
      },
      "tags": [
        "district"
      ],
      "distance": {
        "value": 6562464.289,
        "text": "4068.73 mi"
      }
    },
    {
      "title": {
        "text": "City of Borjomi",
        "hl": [
          {
            "begin": 8,
            "end": 12,
            "type": "MISPRINT"
          }
        ]
      },
      "subtitle": {
        "text": "край Самцхе-Джавахети"
      },
      "tags": [
        "locality"
      ],
      "distance": {
        "value": 6359917.658,
        "text": "3943.15 mi"
      }
    }
  ]
}`
responseSugNotFound = `{
  "suggest_reqid": "1726095099605352-4131673751-suggest-maps-yp-2"
}`
)