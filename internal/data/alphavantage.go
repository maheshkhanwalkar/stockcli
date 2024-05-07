package data

import (
	"errors"
	"stockcli/internal/net"
	"strconv"
	"time"
)

const alphaQuoteUrl = "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol="
const alphaCompanyInfoUrl = "https://www.alphavantage.co/query?function=OVERVIEW&symbol="
const alphaHistoricDataUrl = "https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol="

type QuoteResponse struct {
	Quote struct {
		Price string `json:"05. price"`
	} `json:"Global Quote"`
}

type CompanyInfoResponse struct {
	Name string
}

type HistoricDataResponse struct {
	TimeSeries map[string]map[string]string `json:"Time Series (Daily)"`
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

func (provider AlphaVantageProvider) HistoricData(ticker string) (*HistoricData, error) {
	args := ticker + "&apikey=" + provider.ApiKey

	fullHistoricDataUrl := alphaHistoricDataUrl + args
	historicData := HistoricDataResponse{}

	if err := net.GetJson(fullHistoricDataUrl, &historicData); err != nil {
		return nil, err
	}

	return parseHistoricData(&historicData, ticker)
}

func parseHistoricData(response *HistoricDataResponse, ticker string) (*HistoricData, error) {
	result := make(map[time.Time]float64)

	if len(response.TimeSeries) == 0 {
		return nil, errors.New("no data returned, check API key")
	}

	for timeStr, value := range response.TimeSeries {
		closingPrice, err := strconv.ParseFloat(value["4. close"], 64)

		if err != nil {
			continue
		}

		layout := "2006-01-02"
		tt, err := time.Parse(layout, timeStr)

		if err != nil {
			continue
		}

		result[tt] = closingPrice
	}

	return &HistoricData{Ticker: ticker, Data: result}, nil
}
