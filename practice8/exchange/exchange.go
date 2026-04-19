package exchange

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RateResponse struct {
	Base     string  `json:"base"`
	Target   string  `json:"target"`
	Rate     float64 `json:"rate"`
	ErrorMsg string  `json:"error,omitempty"`
}

type ExchangeService struct {
	BaseURL string
	Client  *http.Client
}

func NewExchangeService(baseURL string) *ExchangeService {
	return &ExchangeService{
		BaseURL: baseURL,
		Client:  &http.Client{Timeout: 5 * time.Second},
	}
}

func (s *ExchangeService) GetRate(from, to string) (float64, error) {
	url := fmt.Sprintf("%s/convert?from=%s&to=%s", s.BaseURL, from, to)

	resp, err := s.Client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result RateResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("api error")
	}

	return result.Rate, nil
}
