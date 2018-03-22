package main

import (
	//	"trading/strategy01"
	"trading/strategy02"
)

func main() {
	// buycoin: "btc"
	// sellcoin: "eur"
	// invest: 1000 (eur)
	// fee: 0.25%
	// period: 60
	// training_iters: 90 minutes
	// win_len_min: 13
	// win_len_max: 34
	//strategy01.Start("btc", "eur", 1000, 30, 90, 11, 24)
	strategy02.Start("btc", "eur", 1000, 0.25, 60, 90, 11, 24, 14, 40.0, 60.0)
}
