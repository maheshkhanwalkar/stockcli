package main

import (
	"flag"
	"os"
)

func parseArgs() string {
	ticker := flag.String("lookup", "", "Stock ticker symbol")
	flag.StringVar(ticker, "l", "", "Stock ticker symbol")

	flag.Parse()

	if *ticker == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return *ticker
}
