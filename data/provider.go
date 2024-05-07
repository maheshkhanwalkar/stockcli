package data

// Quote a stock quote
type Quote struct {
	// Ticker stock ticker
	Ticker string
	// Name company name
	Name string
	// Price quoted stock price
	Price float64
}

type HistoricData struct {
	// Ticker stock ticker
	Ticker string
	// Data historic data keyed by date
	Data map[string]float64
}

// Provider data provider interface
type Provider interface {
	Quote(ticker string) (*Quote, error)
	HistoricData(ticker string) (*HistoricData, error)
}
