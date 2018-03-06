package strategy01

import (
	"fmt"
	"time"

	"github.com/aicastell/cryptotradingbot/src/finantial/ema"
	"github.com/aicastell/cryptotradingbot/src/markets/exchange"
	"github.com/aicastell/cryptotradingbot/src/markets/generic"
)

var UNDEF = int32(-1)
var TRUE = int32(1)
var FALSE = int32(0)

func Start(buycoin string, sellcoin string, invest float64, period time.Duration, training_iters int, win_len_min int, win_len_max int, exchange exchange.Exchange) {
	var emaFast ema.TFinantial_EMA
	var emaSlow ema.TFinantial_EMA
	var emaVol ema.TFinantial_EMA

	emaFast.Reset(win_len_min)
	emaSlow.Reset(win_len_max)
	emaVol.Reset(10)

	fastGtSlow := int32(UNDEF)

	var market generic.TMarket

	market.Reset(buycoin, sellcoin, invest)

	pair := exchange.FormatCoinPair(buycoin, sellcoin) // btceur

	fmt.Println(pair)

	if exchange.HasInmediateTraining() {
		startTraining := time.Now()
		trainingModel(period, pair, &emaFast, &emaSlow, market, exchange)
		endTraining := time.Now()
		fmt.Printf("%v seconds training.\n", endTraining.Sub(startTraining))

		for {
			time.Sleep(period)

			coinPrice, err := exchange.DoGet(pair)
			if err != nil {
				fmt.Printf("Error getting current value %v. %v\n", pair, err)
			}

			updateEma(coinPrice, &emaFast, &emaSlow)
			runEmas(&emaFast, &emaSlow, coinPrice, &fastGtSlow, market)
		}
	} else {
		iter := 0
		for {
			time.Sleep(period)

			coinPrice, err := exchange.DoGet(pair)
			if err != nil {
				fmt.Printf("Error getting current value %v. %v\n", pair, err)
			}

			updateEma(coinPrice, &emaFast, &emaSlow)

			if iter < training_iters {
				iter++
				fmt.Println(iter, training_iters)
				continue
			}

			runEmas(&emaFast, &emaSlow, coinPrice, &fastGtSlow, market)
		}
	}
}

func trainingModel(period time.Duration, pair string, emaFast, emaSlow *ema.TFinantial_EMA, market generic.TMarket, exchange exchange.Exchange) {
	last24hValues, err := exchange.DoGetTrainingValues(period, pair)
	if err != nil {
		fmt.Printf("Error getting current value %v. %v\n", pair, err)
	}

	fmt.Println("Last 24h results: ", last24hValues)

	for _, coinPrice := range last24hValues {
		updateEma(coinPrice, emaFast, emaSlow)
	}
}

func runEmas(emaFast, emaSlow *ema.TFinantial_EMA, coinPrice float64, fastGtSlow *int32, market generic.TMarket) {
	// Initialize fast_gt_slow only once after training
	if *fastGtSlow == UNDEF {
		if emaFast.Ema() > emaSlow.Ema() {
			*fastGtSlow = TRUE
		} else {
			*fastGtSlow = FALSE
		}
		fmt.Println("Training ready. Starting trade now...")
		return
	}

	if *fastGtSlow == FALSE {
		if emaFast.Ema() < emaSlow.Ema() {
			fmt.Println("ema_fast < ema_slow... Se mantiene la tendencia de bajada")
			// tendency is maintained (falling price)
			return
		}
		if market.InsideMarket() == false {
			market.DoBuy(coinPrice)
			fmt.Println("********************************** Buy at: ", market.LastBuyPrice())
			fmt.Println("********************************** CRYPTO: ", market.Crypto())
		} else {
			fmt.Println("===> Tocaba comprar pero ya estoy dentro")
		}
		*fastGtSlow = TRUE
	} else {
		if emaFast.Ema() > emaSlow.Ema() {
			fmt.Println("ema_fast > ema_slow... Se mantiene la tendencia de subida")
			// tendency is maintained (climbing price)
			return
		}
		if market.InsideMarket() == true {
			market.DoSell(coinPrice)
			fmt.Println("********************************** Sell at: ", market.LastSellPrice())
			fmt.Println("********************************** FIAT: ", market.Fiat())
		} else {
			fmt.Println("===> Tocaba vender pero estoy fuera")

		}
		*fastGtSlow = FALSE
	}
}

func updateEma(coinPrice float64, emaFast, emaSlow *ema.TFinantial_EMA) {
	emaFast.New_price(coinPrice)
	emaSlow.New_price(coinPrice)
	fmt.Println("price: ", coinPrice, "ema_fast: ", emaFast.Ema(), "ema_slow: ", emaSlow.Ema(), "time: ", time.Now())
}
