package suggest 
import (
	"fmt"
)
// SuggestResponse represents the full response from the Yandex Suggest API.

const (
    // YandexSuggestAPI is the base URL for the Yandex Suggest API.
    yandexSuggestAPI = "https://suggest-maps.yandex.ru/v1/suggest"
)
type (

Suggest struct {
	lang string `json:"lang"`//ISO639-1
	results int `json:"results"`//Number of results
}

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

func NewSuggest(locale string, resultsNumber int) *Suggest {
	return &Suggest{
		lang:    locale,
        results: resultsNumber,
	}
}

func (s *Suggest) BaseURL(apiKey, text string) string {
	fmt.Println("suggest params is", s.prepareSuggestParams())
    return  fmt.Sprintf("%s?apikey=%s&text=%s&%s", yandexSuggestAPI, apiKey, text, s.prepareSuggestParams())
}

func (s *Suggest) prepareSuggestParams() string {
	return fmt.Sprintf("lang=%s&results=%d",s.lang, s.results)
}