package main

import (
	"fmt"
	"github.com/guptarohit/asciigraph"
	"golang.org/x/exp/maps"
	"log"
	"os"
	"sort"
	"stockcli/config"
	"stockcli/data"
)

func main() {
	args := config.ParseCmdlineArgs()
	configs, err := config.LoadConfigFile()

	if err != nil {
		log.Fatal("Error loading config file: ", err)
	}

	provider := getDataProvider(configs)

	if args.CmdType == config.LOOKUP {
		lookupQuote(args.Ticker, provider)
	} else if args.CmdType == config.GRAPH {
		graphHistoricData(args.Ticker, provider)
	} else {
		log.Fatal("Unknown command: ", args.CmdType)
	}
}

func getDataProvider(configs map[string]string) data.Provider {
	return &data.AlphaVantageProvider{ApiKey: configs["ApiKey"]}
}

func lookupQuote(ticker string, provider data.Provider) {
	quote, err := provider.Quote(ticker)

	if err != nil {
		fmt.Println("Failed to retrieve quote: " + err.Error())
		os.Exit(1)
	}

	fmt.Printf("Stock: %s, Name: %s, Price: %f\n", quote.Ticker, quote.Name, quote.Price)
}

func graphHistoricData(ticker string, provider data.Provider) {
	historicData, err := provider.HistoricData(ticker)

	if err != nil {
		fmt.Println("Failed to retrieve historic data: " + err.Error())
		os.Exit(1)
	}

	graphData := extractGraphData(historicData)
	graph := asciigraph.Plot(graphData, asciigraph.Height(30), asciigraph.Width(80))

	fmt.Println(graph)
}

func extractGraphData(data *data.HistoricData) []float64 {
	result := make([]float64, 0, len(data.Data))

	rawDays := maps.Keys(data.Data)
	sort.Sort(sort.Reverse(sort.StringSlice(rawDays)))

	days := make([]string, 30)
	copy(days, rawDays)
	sort.Strings(days)

	for _, day := range days {
		result = append(result, data.Data[day])
	}

	return result
}
