package main

import (
	//	"trading/strategy01"
	"config"
	"trading/strategy02"
)

func main() {
	var gconf config.TConfig_Bot

	gconf.LoadConfig()
	gconf.Log()

	strategy02.Start(gconf.Global.BuyCoin,
		gconf.Global.SellCoin,
		gconf.Global.Invest,
		gconf.Global.Fee,
		gconf.Global.Period,
		gconf.Global.TrainingIters,
		gconf.EMA.Fast,
		gconf.EMA.Slow,
		gconf.RSI.WinLen,
		gconf.RSI.BuyLevel,
		gconf.RSI.SellLevel)
}
