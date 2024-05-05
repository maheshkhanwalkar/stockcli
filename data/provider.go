package data

// Provider data provider interface
type Provider interface {
	Quote(ticker string) (float64, error)
}
