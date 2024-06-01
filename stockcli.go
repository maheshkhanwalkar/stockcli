package main

import (
	"fmt"
	"github.com/guptarohit/asciigraph"
	"github.com/maheshkhanwalkar/stockcli/internal/config"
	"github.com/maheshkhanwalkar/stockcli/internal/data"
	"golang.org/x/exp/maps"
	"golang.org/x/term"
	"log"
	"os"
	"sort"
	"syscall"
	"time"
)

func main() {
	args := config.ParseCmdlineArgs()

	if args.CmdType == config.INIT {
		setupConfiguration()
		return
	}

	configs, err := config.LoadConfigFile()

	if err != nil {
		log.Fatal("Error loading config file: ", err)
	}

	provider := getDataProvider(configs)

	if args.CmdType == config.LOOKUP {
		lookupQuote(args.Ticker, provider)
	} else if args.CmdType == config.GRAPH {
		graphHistoricData(args.Ticker, provider, args.Days)
	} else {
		log.Fatal("Unknown command: ", args.CmdType)
	}
}

func getDataProvider(configs map[string]string) data.Provider {
	return &data.AlphaVantageProvider{ApiKey: configs["ApiKey"]}
}

func setupConfiguration() {
	println("Enter API key: ")
	key, err := term.ReadPassword(syscall.Stdin)

	if err != nil {
		log.Fatal("Failed to read API key", err)
	}

	configs := make(map[string]string)
	configs["ApiKey"] = string(key)

	err = config.CreateConfigFile(configs)
	if err != nil {
		log.Fatal("Failed to create config file:", err)
	}

	log.Println("Successfully initialized stockcli")
}

func lookupQuote(ticker string, provider data.Provider) {
	quote, err := provider.Quote(ticker)

	if err != nil {
		fmt.Println("Failed to retrieve quote:", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Stock: %s, Name: %s, Price: %f\n", quote.Ticker, quote.Name, quote.Price)
}

func graphHistoricData(ticker string, provider data.Provider, days int) {
	historicData, err := provider.HistoricData(ticker)

	if err != nil {
		fmt.Println("Failed to retrieve historic data:", err.Error())
		os.Exit(1)
	}

	graphData := extractGraphData(historicData, days)
	colour := graphColour(graphData)

	graph := asciigraph.Plot(graphData, asciigraph.Height(30), asciigraph.Width(80),
		asciigraph.SeriesColors(colour))

	fmt.Println(graph)
}

func extractGraphData(data *data.HistoricData, numDays int) []float64 {
	result := make([]float64, 0, len(data.Data))
	rawDays := maps.Keys(data.Data)

	sort.Slice(rawDays, func(i, j int) bool {
		return rawDays[i].After(rawDays[j])
	})

	days := make([]time.Time, min(numDays, len(data.Data)))
	copy(days, rawDays)
	sort.Slice(days, func(i, j int) bool {
		return days[i].Before(days[j])
	})

	for _, day := range days {
		result = append(result, data.Data[day])
	}

	return result
}

func graphColour(graphData []float64) asciigraph.AnsiColor {
	if len(graphData) == 0 {
		return asciigraph.Default
	}

	first := graphData[0]
	last := graphData[len(graphData)-1]

	if first < last {
		return asciigraph.Green
	} else if first > last {
		return asciigraph.Red
	} else {
		return asciigraph.Default
	}
}
