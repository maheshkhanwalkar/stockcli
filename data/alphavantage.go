package data

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

const QuoteUrl = "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol="

type QuoteResponse struct {
	Quote struct {
		Price string `json:"05. price"`
	} `json:"Global Quote"`
}

type AlphaVantageProvider struct {
	ApiKey string
}

func (provider AlphaVantageProvider) Quote(ticker string) (float64, error) {
	fullUrl := QuoteUrl + ticker + "&apikey=" + provider.ApiKey
	resp, err := http.Get(fullUrl)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return 0, err
	}

	quoteResponse := QuoteResponse{}

	if err = json.Unmarshal(body, &quoteResponse); err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(quoteResponse.Quote.Price, 64)
	return price, err
}
