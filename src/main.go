package main

import (
	"time"

	"github.com/aicastell/cryptotradingbot/src/trading/strategy01"
	"github.com/carlesar/cryptotradingbot/src/markets/poloniex"
)

func main() {
	// buycoin: "btc"
	// sellcoin: "eur"
	// invest: 1000 (eur)
	// period: 60
	// training_iters: 90 minutes
	// win_len_min: 13
	// win_len_max: 34
	// exchange: poloniex
	strategy01.Start("BTC", "XMR", 1000, 300*time.Second, 90, 13, 34, &poloniex.Poloniex{})
	//strategy01.Start("btc", "eur", 1000, 10*time.Second, 90, 13, 34, &bitstamp.Bitstamp{})
}
