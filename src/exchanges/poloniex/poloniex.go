package poloniex

import (
    "config"
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"strconv"
)

type TPoloniex struct{
    Config     *config.TBotConfig
}

func (ob *TPoloniex) SetConfig(cfg *config.TBotConfig) {
    ob.Config = cfg
}

func (ob *TPoloniex) GetPrice(coinpair string) (lastValue float64, err error) {
/*
    url := "https://poloniex.com/public?command=returnTicker"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("ERROR creating http tickers request.", err)
        return -1, &utils.MyError{13}
    }
    client := http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("ERROR making http tockers request.", err)
        return -1, &utils.MyError{13}
    }
    defer resp.Body.Close()

    var allTickers map[string]map[string]interface{}
    err = json.NewDecoder(resp.Body).Decode(&allTickers)
    if err != nil {
        fmt.Println("ERROR decoding tickers response.", err)
        return -1, &utils.MyError{13}
    }

    for ticker, content := range allTickers {
        if ticker == coinpair {
            lastValue, err := ei.N(content).M("last").Float64()
            if err != nil {
                fmt.Println("ERROR parsing coinpair last value to float64.", err)
                return -1, &utils.MyError{13}
            }
            return lastValue, nil
        }
    }
    */
    return 0, nil
}

