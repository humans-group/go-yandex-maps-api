package client

type (
	SuggestResponse struct {
		SuggestReqID string          `json:"suggest_reqid"`
		Results      []SuggestResult `json:"results"`
	}

	SuggestResult struct {
		Title    Title    `json:"title"`
		Subtitle Subtitle `json:"subtitle,omitempty"`
		Tags     []string `json:"tags"`
		Distance Distance `json:"distance"`
	}

	Title struct {
		Text string    `json:"text"`
		HL   []HLRange `json:"hl,omitempty"`
	}

	HLRange struct {
		Begin int `json:"begin"`
		End   int `json:"end"`
	}

	Subtitle struct {
		Text string `json:"text"`
	}

	Distance struct {
		Value float64 `json:"value"`
		Text  string  `json:"text"`
	}
	//Geocode response from Yandex GeoCoder API docs and https://github.com/codingsince1985/geo-golang project
	GeocodeResponse struct {
		Response struct {
			GeoObjectCollection struct {
				MetaDataProperty struct {
					GeocoderResponseMetaData struct {
						Request string `json:"request"`
						Found   string `json:"found"`
						Results string `json:"results"`
					} `json:"GeocoderResponseMetaData"`
				} `json:"metaDataProperty"`
				FeatureMember []*YandexFeatureMember `json:"featureMember"`
			} `json:"GeoObjectCollection"`
		} `json:"response"`
	}
	YandexFeatureMember struct {
		GeoObject struct {
			MetaDataProperty struct {
				GeocoderMetaData struct {
					Kind      string `json:"kind"`
					Text      string `json:"text"`
					Precision string `json:"precision"`
					Address   struct {
						CountryCode string `json:"country_code"`
						PostalCode  string `json:"postal_code"`
						Formatted   string `json:"formatted"`
						Components  []struct {
							Kind string `json:"kind"`
							Name string `json:"name"`
						} `json:"Components"`
					} `json:"Address"`
				} `json:"GeocoderMetaData"`
			} `json:"metaDataProperty"`
			Description string `json:"description"`
			Name        string `json:"name"`
			BoundedBy   struct {
				Envelope struct {
					LowerCorner string `json:"lowerCorner"`
					UpperCorner string `json:"upperCorner"`
				} `json:"Envelope"`
			} `json:"boundedBy"`
			Point struct {
				Pos string `json:"pos"`
			} `json:"Point"`
		} `json:"GeoObject"`
	}
)
