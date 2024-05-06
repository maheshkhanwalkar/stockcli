package data

import (
	"errors"
	"stockcli/net"
	"strconv"
)

const alphaQuoteUrl = "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol="
const alphaCompanyInfoUrl = "https://www.alphavantage.co/query?function=OVERVIEW&symbol="

type QuoteResponse struct {
	Quote struct {
		Price string `json:"05. price"`
	} `json:"Global Quote"`
}

type CompanyInfoResponse struct {
	Name string
}

// AlphaVantageProvider data provider
type AlphaVantageProvider struct {
	ApiKey string
}

func (provider AlphaVantageProvider) Quote(ticker string) (*Quote, error) {
	args := ticker + "&apikey=" + provider.ApiKey

	fullQuoteUrl := alphaQuoteUrl + args
	fullCompanyInfoUrl := alphaCompanyInfoUrl + args

	quoteResponse := QuoteResponse{}
	companyInfo := CompanyInfoResponse{}

	if err := net.GetJson(fullQuoteUrl, &quoteResponse); err != nil {
		return nil, err
	}

	if err := net.GetJson(fullCompanyInfoUrl, &companyInfo); err != nil {
		return nil, err
	}

	if quoteResponse.Quote.Price == "" {
		return nil, errors.New(ticker + " does not exist")
	}

	price, err := strconv.ParseFloat(quoteResponse.Quote.Price, 64)

	if err != nil {
		return nil, err
	}

	return &Quote{Ticker: ticker, Price: price, Name: companyInfo.Name}, err
}
