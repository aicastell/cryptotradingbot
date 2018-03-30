package main

import (
	"config"
	"exchanges/bitstamp"
	"exchanges/poloniex"
	"fetcher"
	"fmt"
	"inputs/datafile"
	"inputs/stdinput"
	"os"
	"trading/strategy01"
	"trading/strategy02"
)

func main() {
	// Read and load configuration file
	var cfg config.TBotConfig
	cfg.LoadConfig()
	cfg.Log()

	// Declare infetch generic pointer
	var infetch *fetcher.TFetcher
	infetch = new(fetcher.TFetcher)

	// Declare different sources of prices
	var src00 bitstamp.TBitstamp
	var src01 poloniex.TPoloniex
	var src02 datafile.TDataFile
	var src03 stdinput.TStdInput

	// Point to proper fetcher
	if cfg.Global.Fetcher == "bitstamp" {
		*infetch = &src00
	} else if cfg.Global.Fetcher == "poloniex" {
		*infetch = &src01
	} else if cfg.Global.Fetcher == "datafile" {
		*infetch = &src02
	} else if cfg.Global.Fetcher == "stdin" {
		*infetch = &src03
	} else {
		fmt.Println("Unknown strategy")
		os.Exit(1)
	}

	// Configure input fetcher
	(*infetch).SetConfig(&cfg)

	// Declare different strategies
	var strategy01 strategy01.TStrag01
	var strategy02 strategy02.TStrag02

	// Start proper strategy
	if cfg.Global.Strategy == "strategy01" {
		strategy01.Start(&cfg, infetch)
	} else if cfg.Global.Strategy == "strategy02" {
		strategy02.Start(&cfg, infetch)
	} else {
		fmt.Println("Unknown strategy")
		os.Exit(1)
	}

	os.Exit(0)
}
