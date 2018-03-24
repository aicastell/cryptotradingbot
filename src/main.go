package main

import (
	//	"trading/strategy01"
    "config"
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

    var gconf config.TConfig_Bot

    gconf.LoadConfig()
    gconf.Log()

	strategy02.Start(gconf.Global.BuyCoin,
                     gconf.Global.SellCoin,
                     gconf.Global.Invest,
                     gconf.Global.Fee,
                     gconf.Global.Period,
                     gconf.Global.TrainingIters,
                     gconf.EMA.WinLenMin,
                     gconf.EMA.WinLenMax,
                     gconf.RSI.WinLen,
                     gconf.RSI.BuyLevel,
                     gconf.RSI.SellLevel)
}
