package data

import "time"

// Quote represents a stock quote
type Quote struct {
	// Ticker stock ticker
	Ticker string
	// Name company name
	Name string
	// Price quoted stock price
	Price float64
}

// HistoricData for a stock
type HistoricData struct {
	// Ticker stock ticker
	Ticker string
	// Data historic data keyed by date
	Data map[time.Time]float64
}

// Provider data provider interface
type Provider interface {
	Quote(ticker string) (*Quote, error)
	HistoricData(ticker string) (*HistoricData, error)
}
