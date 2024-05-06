package data

import (
	"errors"
	"stockcli/net"
	"strconv"
)

const alphaQuoteUrl = "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol="

type QuoteResponse struct {
	Quote struct {
		Price string `json:"05. price"`
	} `json:"Global Quote"`
}

// AlphaVantageProvider data provider
type AlphaVantageProvider struct {
	ApiKey string
}

func (provider AlphaVantageProvider) Quote(ticker string) (*Quote, error) {
	fullUrl := alphaQuoteUrl + ticker + "&apikey=" + provider.ApiKey
	quoteResponse := QuoteResponse{}

	if err := net.GetJson(fullUrl, &quoteResponse); err != nil {
		return nil, err
	}

	if quoteResponse.Quote.Price == "" {
		return nil, errors.New(ticker + " does not exist")
	}

	price, err := strconv.ParseFloat(quoteResponse.Quote.Price, 64)
	return &Quote{Ticker: ticker, Price: price}, err
}
