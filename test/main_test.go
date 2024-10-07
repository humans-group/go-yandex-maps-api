package main_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/humans-group/go-yandex-maps-api/services/geocode"
	client "github.com/humans-group/go-yandex-maps-api/services/httpclient"
	"github.com/humans-group/go-yandex-maps-api/services/suggest"
	"github.com/stretchr/testify/assert"
)

var (
	tokenSuggest = os.Getenv("YANDEX_SUGGEST_API_KEY")
	tokenGeocode = os.Getenv("YANDEX_GEOCODE_API_KEY")
)

func testServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(response))
	}))
}

func testServerWithDelay(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		time.Sleep(200 * time.Millisecond) // simulate network delay
		resp.Write([]byte(response))
	}))
}

func TestYandexGeoSuggest(t *testing.T) {
	ts := testServer(responseSugFound)
	defer ts.Close()
	clientAPI := &client.SimpleHTTPClient{}
	suggestAPI := suggest.NewSuggestAPI(tokenSuggest, ts.URL+"/")
	suggestAPI.AddLanguage("eng")
	suggestion, err := client.Suggest(clientAPI, suggestAPI, "Burj")

	assert.NoError(t, err)
	assert.True(t, len(suggestion.Results) > 0)
	assert.True(t, suggestion.Results[0].Title.Text == "Burji")
	assert.True(t, suggestion.Results[1].Title.Text == "Downtown Dubai")
	assert.True(t, suggestion.Results[2].Title.Text == "City of Borjomi")
}

func TestYandexGeoSuggestNoResult(t *testing.T) {
	ts := testServer(responseSugNotFound)
	defer ts.Close()
	clientAPI := &client.SimpleHTTPClient{}
	suggestAPI := suggest.NewSuggestAPI(tokenSuggest, ts.URL+"/")
	suggestAPI.AddLanguage("eng")
	suggestion, err := client.Suggest(clientAPI, suggestAPI, "Burj")
	assert.Nil(t, err)
	assert.Nil(t, suggestion.Results)
}

func TestYandexForwardGeoCodeResult(t *testing.T) {
	ts := testServer(responseGeocodeFound)
	defer ts.Close()
	clientAPI := &client.SimpleHTTPClient{}
	geocodeAPI := geocode.NewGeocodeAPI(tokenGeocode, ts.URL+"/")
	geocodeAPI.AddLanguage("en_US")
	_, err := client.ForwardGeocode(clientAPI, geocodeAPI, "Mohammed Bin Rashid Boulevard 1")
	assert.Nil(t, err)
}

func TestYandexReverseGeoCodeResult(t *testing.T) {
	ts := testServer(responseGeocodeFound)
	defer ts.Close()
	clientAPI := &client.SimpleHTTPClient{}
	geocodeAPI := geocode.NewGeocodeAPI(tokenGeocode, ts.URL+"/")
	geocodeAPI.AddLanguage("en_US")
	_, err := client.ReverseGeocode(clientAPI, geocodeAPI, 69.02, 42.01)
	assert.Nil(t, err)
}

func TestYandexForwardGeoCodeNoResult(t *testing.T) {
	ts := testServer(responseGeocodeNotFound)
	defer ts.Close()
	clientAPI := &client.SimpleHTTPClient{}
	geocodeAPI := geocode.NewGeocodeAPI(tokenGeocode, ts.URL+"/")
	geocodeAPI.AddLanguage("en_US")
	res, err := client.ForwardGeocode(clientAPI, geocodeAPI, "{dsaffdsa}")
	assert.Nil(t, err)
	assert.Equal(t, len(res.Response.GeoObjectCollection.FeatureMember), 0)
}

func TestYandexServerWithTimeout(t *testing.T) {
	ts := testServerWithDelay(responseGeocodeNotFound)
	defer ts.Close()
	clientAPI := &client.SimpleHTTPClient{Timeout: 100 * time.Millisecond}
	geocodeAPI := geocode.NewGeocodeAPI(tokenGeocode, ts.URL+"/")
	geocodeAPI.AddLanguage("en_US")
	_, err := client.ForwardGeocode(clientAPI, geocodeAPI, "Mohammed Bin Rashid Boulevard 1")
	if assert.Error(t, err) {
		assert.Equal(t, err, client.ErrTimeout)
	}

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
	responseGeocodeFound = `{
  "response": {
    "GeoObjectCollection": {
      "metaDataProperty": {
        "GeocoderResponseMetaData": {
          "request": "Mohammed Bin Rashid Boulevard 1",
          "results": "10",
          "found": "2"
        }
      },
      "featureMember": [
        {
          "GeoObject": {
            "metaDataProperty": {
              "GeocoderMetaData": {
                "precision": "exact",
                "text": "1, Mohammed Bin Rashid Boulevard, Downtown Dubai, Dubai, United Arab Emirates",
                "kind": "house",
                "Address": {
                  "country_code": "AE",
                  "formatted": "1, Mohammed Bin Rashid Boulevard, Downtown Dubai, Dubai, United Arab Emirates",
                  "Components": [
                    {
                      "kind": "country",
                      "name": "United Arab Emirates"
                    },
                    {
                      "kind": "province",
                      "name": "Dubai"
                    },
                    {
                      "kind": "area",
                      "name": "Sector 3"
                    },
                    {
                      "kind": "district",
                      "name": "Downtown Dubai"
                    },
                    {
                      "kind": "district",
                      "name": "Downtown Dubai"
                    },
                    {
                      "kind": "street",
                      "name": "Mohammed Bin Rashid Boulevard"
                    },
                    {
                      "kind": "house",
                      "name": "1"
                    }
                  ]
                },
                "AddressDetails": {
                  "Country": {
                    "AddressLine": "1, Mohammed Bin Rashid Boulevard, Downtown Dubai, Dubai, United Arab Emirates",
                    "CountryNameCode": "AE",
                    "CountryName": "United Arab Emirates",
                    "AdministrativeArea": {
                      "AdministrativeAreaName": "Dubai",
                      "SubAdministrativeArea": {
                        "SubAdministrativeAreaName": "Sector 3",
                        "Locality": {
                          "DependentLocality": {
                            "DependentLocalityName": "Downtown Dubai",
                            "DependentLocality": {
                              "DependentLocalityName": "Downtown Dubai",
                              "Thoroughfare": {
                                "ThoroughfareName": "Mohammed Bin Rashid Boulevard",
                                "Premise": {
                                  "PremiseNumber": "1"
                                }
                              }
                            }
                          }
                        }
                      }
                    }
                  }
                }
              }
            },
            "name": "1, Mohammed Bin Rashid Boulevard",
            "description": "Downtown Dubai, Dubai, United Arab Emirates",
            "boundedBy": {
              "Envelope": {
                "lowerCorner": "55.270141 25.193445",
                "upperCorner": "55.278352 25.200915"
              }
            },
            "uri": "ymapsbm1://geo?data=CgoyMTE3NTQxODgxEocB2KfZhNil2YXYp9ix2KfYqiDYp9mE2LnYsdio2YrYqSDYp9mE2YXYqtit2K_YqSwg2KXZhdin2LHYqSDYr9io2YosINmI2LPYtyDZhdiv2YrZhtipINiv2KjZiiwg2KjZiNmE2YrZgdin2LHYryDZhdit2YXYryDYqNmGINix2KfYtNivLCAxIgoN1BhdQhXVk8lB",
            "Point": {
              "pos": "55.274247 25.19718"
            }
          }
        },
        {
          "GeoObject": {
            "metaDataProperty": {
              "GeocoderMetaData": {
                "precision": "street",
                "text": "Mohammed Bin Rashid Boulevard, Downtown Dubai, Dubai, United Arab Emirates",
                "kind": "street",
                "Address": {
                  "country_code": "AE",
                  "formatted": "Mohammed Bin Rashid Boulevard, Downtown Dubai, Dubai, United Arab Emirates",
                  "Components": [
                    {
                      "kind": "country",
                      "name": "United Arab Emirates"
                    },
                    {
                      "kind": "province",
                      "name": "Dubai"
                    },
                    {
                      "kind": "area",
                      "name": "Sector 3"
                    },
                    {
                      "kind": "district",
                      "name": "Downtown Dubai"
                    },
                    {
                      "kind": "street",
                      "name": "Mohammed Bin Rashid Boulevard"
                    }
                  ]
                },
                "AddressDetails": {
                  "Country": {
                    "AddressLine": "Mohammed Bin Rashid Boulevard, Downtown Dubai, Dubai, United Arab Emirates",
                    "CountryNameCode": "AE",
                    "CountryName": "United Arab Emirates",
                    "AdministrativeArea": {
                      "AdministrativeAreaName": "Dubai",
                      "SubAdministrativeArea": {
                        "SubAdministrativeAreaName": "Sector 3",
                        "Locality": {
                          "DependentLocality": {
                            "DependentLocalityName": "Downtown Dubai",
                            "Thoroughfare": {
                              "ThoroughfareName": "Mohammed Bin Rashid Boulevard"
                            }
                          }
                        }
                      }
                    }
                  }
                }
              }
            },
            "name": "Mohammed Bin Rashid Boulevard",
            "description": "Downtown Dubai, Dubai, United Arab Emirates",
            "boundedBy": {
              "Envelope": {
                "lowerCorner": "55.277606 25.195668",
                "upperCorner": "55.279807 25.199027"
              }
            },
            "uri": "ymapsbm1://geo?data=Cgo1ODA2NzkyODQ1EkpVbml0ZWQgQXJhYiBFbWlyYXRlcywgRHViYWksIERvd250b3duIER1YmFpLCBNb2hhbW1lZCBCaW4gUmFzaGlkIEJvdWxldmFyZCIKDXQdXUIVGZTJQQ,,",
            "Point": {
              "pos": "55.278765 25.197311"
            }
          }
        }
      ]
    }
  }
}`
	responseGeocodeNotFound = `{
  "response":{
  "GeoObjectCollection":{
  "metaDataProperty":{
  "GeocoderResponseMetaData":{
  "request":" -*-/","results":"10","found":"0"}
  },"featureMember":[]}
  }
}`
)
