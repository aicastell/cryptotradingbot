package poloniex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aicastell/cryptotradingbot/src/utils"
	"github.com/jaracil/ei"
)

var validPeriods = []float64{300, 900, 1800, 7200, 14400, 86400}
var hasInmediateTraining bool = true

type Poloniex struct{}

func (Poloniex) DoGet(coinpair string) (lastValue float64, err error) {
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
	return 0, nil
}

func (Poloniex) DoGetTrainingValues(period time.Duration, coinpair string) (last24hValues []float64, err error) {
	periodSeconds := period.Seconds()
	if !utils.ContainsFloat64(validPeriods, periodSeconds) {
		fmt.Printf("ERROR invalid period %d. Expected: %v.\n", periodSeconds, validPeriods)
		return nil, &utils.MyError{13}
	}
	start := time.Now().Add(-24 * time.Hour).Unix()
	url := fmt.Sprintf("https://poloniex.com/public?command=returnChartData&currencyPair=%s&start=%d&end=9999999999&period=%v", coinpair, start, periodSeconds)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("ERROR creating http last24h values request.", err)
		return nil, &utils.MyError{13}
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR making http tockers request.", err)
		return nil, &utils.MyError{13}
	}
	defer resp.Body.Close()

	var values []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&values)
	if err != nil {
		fmt.Println("ERROR decoding last24h values response %v.", values, err)
		return nil, &utils.MyError{13}
	}
	fmt.Println(values)
	parsedValues := []float64{}
	for _, value := range values {
		valueF, err := ei.N(value).M("close").Float64()
		if err != nil {
			fmt.Println("ERROR parsing value to float64.", err)
			return nil, &utils.MyError{13}
		}
		parsedValues = append(parsedValues, valueF)
	}

	return parsedValues, nil
}

func (Poloniex) FormatCoinPair(buyCoin, sellCoin string) string {
	return strings.ToUpper(buyCoin) + "_" + strings.ToUpper(sellCoin)
}

func (Poloniex) HasInmediateTraining() bool {
	return true
}
