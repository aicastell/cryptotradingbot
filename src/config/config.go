package config

/*
{
    "global": {
        "strategy": "strategy02",
        "fetcher": "datafile",
        "buycoin": "btc",
        "sellcoin": "eur",
        "invest": 1000,
        "fee": 0.25,
        "period": 0,
        "training_iters": 90
    },
    "strategy01": {
        "name": "Pure EMA",
        "ema": {
            "fast": 11,
            "slow": 24
        }
    },
    "strategy02": {
        "name": "Ema + RSI",
        "ema": {
            "fast": 11,
            "slow": 24
        },
        "rsi": {
            "win_len": 14,
            "buy_level": 40.0,
            "sell_level": 60.0
        }
    },
    "strategy03": {
        "name": "Pure MACD",
        "ema": {
            "fast": 11,
            "slow": 24
        }
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

type TConfig_Strat01 struct {
	Name string      `json:"name"`
	EMA  TConfig_EMA `json:"ema"`
}

type TConfig_Strat02 struct {
	Name string      `json:"name"`
	EMA  TConfig_EMA `json:"ema"`
	RSI  TConfig_RSI `json:"rsi"`
}

type TConfig_Strat03 struct {
	Name string      `json:"name"`
	EMA  TConfig_EMA `json:"ema"`
	// TODO
}

type TBotConfig struct {
	Global  TConfig_Global  `json:"global"`
	Strat01 TConfig_Strat01 `json:"strategy01"`
	Strat02 TConfig_Strat02 `json:"strategy02"`
	Strat03 TConfig_Strat03 `json:"strategy03"`
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
	fmt.Println("Global configuration:")
	fmt.Println("**********************************")
	fmt.Println("Strategy:\t", gconf.Global.Strategy)
	fmt.Println("Fetcher:\t", gconf.Global.Fetcher)
	fmt.Println("BuyCoin:\t", gconf.Global.BuyCoin)
	fmt.Println("SellCoin:\t", gconf.Global.SellCoin)
	fmt.Println("Invest:\t\t", gconf.Global.Invest)
	fmt.Println("Fee:\t\t", gconf.Global.Fee)
	fmt.Println("Period:\t\t", gconf.Global.Period)
	fmt.Println("TrainingIters:\t", gconf.Global.TrainingIters)

	fmt.Println("**********************************")
	fmt.Println("Strategy01 summary:")
	fmt.Println("**********************************")
	fmt.Println("Strat01 name: \t", gconf.Strat01.Name)
	fmt.Println("Strat01 ema fast: \t", gconf.Strat01.EMA.Fast)
	fmt.Println("Strat01 ema slow: \t", gconf.Strat01.EMA.Slow)

	fmt.Println("**********************************")
	fmt.Println("Strategy02 summary:")
	fmt.Println("**********************************")
	fmt.Println("Strat02 name: \t", gconf.Strat02.Name)
	fmt.Println("Strat02 ema fast: \t", gconf.Strat02.EMA.Fast)
	fmt.Println("Strat02 ema slow: \t", gconf.Strat02.EMA.Slow)
	fmt.Println("Strat02 rsi window len: \t", gconf.Strat02.RSI.WinLen)
	fmt.Println("Strat02 rsi buy level: \t", gconf.Strat02.RSI.BuyLevel)
	fmt.Println("Strat02 rsi sell level: \t", gconf.Strat02.RSI.SellLevel)

	fmt.Println("**********************************")
	fmt.Println("Strategy03 summary:")
	fmt.Println("**********************************")
	fmt.Println("Strat03 name: \t", gconf.Strat03.Name)
	fmt.Println("Strat03 ema fast: \t", gconf.Strat03.EMA.Fast)
	fmt.Println("Strat03 ema slow: \t", gconf.Strat03.EMA.Slow)
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
