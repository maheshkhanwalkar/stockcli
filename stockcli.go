package main

import (
	"fmt"
	"os"
	"stockcli/data"
)

func main() {
	ticker := parseArgs()
	println("Looking up basic stock data for " + ticker)

	provider := getDataProvider()
	price, err := provider.Quote(ticker)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Stock: %s, Price: %f\n", ticker, price)
}

func getDataProvider() data.Provider {
	// TODO - get from config
	return &data.AlphaVantageProvider{ApiKey: "REDACTED"}
}
