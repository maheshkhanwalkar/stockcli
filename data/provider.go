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

// Provider data provider interface
type Provider interface {
	Quote(ticker string) (*Quote, error)
}
