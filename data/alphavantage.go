package data

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

const quoteUrl = "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol="

type QuoteResponse struct {
	Quote struct {
		Price string `json:"05. price"`
	} `json:"Global Quote"`
}

type AlphaVantageProvider struct {
	ApiKey string
}

func (provider AlphaVantageProvider) Quote(ticker string) (*Quote, error) {
	fullUrl := quoteUrl + ticker + "&apikey=" + provider.ApiKey
	resp, err := http.Get(fullUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	quoteResponse := QuoteResponse{}

	if err = json.Unmarshal(body, &quoteResponse); err != nil {
		return nil, err
	}

	if quoteResponse.Quote.Price == "" {
		return nil, errors.New(ticker + " does not exist")
	}

	price, err := strconv.ParseFloat(quoteResponse.Quote.Price, 64)
	return &Quote{Ticker: ticker, Price: price}, err
}
