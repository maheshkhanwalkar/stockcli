package main

import (
	"fmt"
	"log"
	"os"
	"stockcli/config"
	"stockcli/data"
)

func main() {
	ticker := config.ParseCmdlineArgs()
	configs, err := config.LoadConfigFile()

	if err != nil {
		log.Fatal("Error loading config file: ", err)
	}

	println("Looking up basic stock data for " + ticker)

	provider := getDataProvider(configs)
	price, err := provider.Quote(ticker)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Stock: %s, Price: %f\n", ticker, price)
}

func getDataProvider(configs map[string]string) data.Provider {
	return &data.AlphaVantageProvider{ApiKey: configs["ApiKey"]}
}
