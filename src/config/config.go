package config

/*
{
    "global": {
        "strategy": "strategy02",
        "fetcher": "poloniex",
        "buycoin": "btc",
        "sellcoin": "eur",
        "invest": 1000,
        "fee": 0.25,
        "period": 60,
        "training_iters": 90
    },
    "ema": {
        "fast": 13,
        "slow": 34
    },
    "rsi": {
        "win_len": 14,
        "buy_level": 40,
        "sell_level": 60
    }
}

*/

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type TConfig_Global struct {
	Strategy      string  `json:"strategy"`
	Fetcher       string  `json:"fetcher"`
	BuyCoin       string  `json:"buycoin"`
	SellCoin      string  `json:"sellcoin"`
	Invest        float64 `json:"invest"`
	Fee           float64 `json:"fee"`
	Period        int     `json:"period"`
	TrainingIters int     `json:"training_iters"`
}

type TConfig_EMA struct {
	Fast int `json:"fast"`
	Slow int `json:"slow"`
}

type TConfig_RSI struct {
	WinLen    int     `json:"win_len"`
	BuyLevel  float64 `json:"buy_level"`
	SellLevel float64 `json:"sell_level"`
}

type TBotConfig struct {
	Global TConfig_Global `json:"global"`
	EMA    TConfig_EMA    `json:"ema"`
	RSI    TConfig_RSI    `json:"rsi"`
}

func (gconf *TBotConfig) LoadConfig() {
	filename := flag.String("c", "/etc/trabot.d/trabot.conf", "Trading bot configuration file")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println("Error1")
		os.Exit(1)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(gconf)
	if err != nil {
		fmt.Println("Error2", err)
		os.Exit(1)
	}
}

func (gconf *TBotConfig) Log() {
	fmt.Println("**********************************")
	fmt.Println("Configuration summary:")
	fmt.Println("**********************************")
	fmt.Println("Strategy:\t", gconf.Global.Strategy)
	fmt.Println("Fetcher:\t", gconf.Global.Fetcher)
	fmt.Println("BuyCoin:\t", gconf.Global.BuyCoin)
	fmt.Println("SellCoin:\t", gconf.Global.SellCoin)
	fmt.Println("Invest:\t\t", gconf.Global.Invest)
	fmt.Println("Fee:\t\t", gconf.Global.Fee)
	fmt.Println("Period:\t\t", gconf.Global.Period)
	fmt.Println("TrainingIters:\t", gconf.Global.TrainingIters)

	fmt.Println("EMA Fast:\t", gconf.EMA.Fast)
	fmt.Println("EMA Slow:\t", gconf.EMA.Slow)

	fmt.Println("RSI WinLen:\t", gconf.RSI.WinLen)
	fmt.Println("RSI BuyLevel:\t", gconf.RSI.BuyLevel)
	fmt.Println("RSI SellLevel:\t", gconf.RSI.SellLevel)
	fmt.Println("**********************************")
}

/*
func main() {

    var gconf TBotConfig

    filename := flag.String("c", "/etc/trabot.d/trabot.conf", "Trading bot configuration file")
    flag.Parse()

    file, err := os.Open(*filename)
    if err != nil {
        fmt.Println("Error1")
        os.Exit(1)
    }

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&gconf)
    if err != nil {
        fmt.Println("Error2", err)
        os.Exit(1)
    }

    fmt.Println(gconf.Global.Strategy)
    fmt.Println(gconf.Global.Fetcher)
    fmt.Println(gconf.Global.BuyCoin)
    fmt.Println(gconf.Global.SellCoin)
    fmt.Println(gconf.Global.Invest)
    fmt.Println(gconf.Global.Fee)
    fmt.Println(gconf.Global.Period)
    fmt.Println(gconf.Global.TrainingIters)

    fmt.Println(gconf.EMA.Fast)
    fmt.Println(gconf.EMA.Slow)

    fmt.Println(gconf.RSI.WinLen)
    fmt.Println(gconf.RSI.BuyLevel)
    fmt.Println(gconf.RSI.SellLevel)
}

*/
