package client

type (
	SuggestResponse struct {
		SuggestReqID string         `json:"suggest_reqid"`
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
)