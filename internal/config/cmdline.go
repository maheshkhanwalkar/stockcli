package config

import (
	"flag"
	"os"
)

type CmdType int

const (
	LOOKUP CmdType = iota
	GRAPH
	INIT
)

type CmdlineArgs struct {
	Ticker  string
	CmdType CmdType
}

func ParseCmdlineArgs() *CmdlineArgs {
	init := flag.Bool("init", false, "initialize stockcli configuration")

	lookupTicker := flag.String("lookup", "", "Look up a stock ticker symbol")
	flag.StringVar(lookupTicker, "l", "", "Look up a stock ticker symbol")

	graphTicker := flag.String("graph", "", "Graph historic data for a stock ticker symbol")
	flag.StringVar(graphTicker, "g", "", "Graph historic data for a stock ticker symbol")

	flag.Parse()

	if *init {
		return &CmdlineArgs{CmdType: INIT}
	}

	if *graphTicker == "" && *lookupTicker == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *graphTicker != "" {
		return &CmdlineArgs{Ticker: *graphTicker, CmdType: GRAPH}
	} else {
		return &CmdlineArgs{Ticker: *lookupTicker, CmdType: LOOKUP}
	}
}
