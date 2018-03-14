package strategy02

import (
	"finantial/ema"
	"finantial/rsi"
	"fmt"
	"markets/bitstamp"
	"markets/generic"
	"time"
)

var UNDEF = int32(-1)
var TRUE = int32(1)
var FALSE = int32(0)

func Start(buycoin string, sellcoin string, invest float64, period time.Duration, training_iters int, win_len_min int, win_len_max int, rsi_win_len int, rsi_buy_level float64, rsi_sell_level float64) {
	var ema_fast ema.TFinantial_EMA
	var ema_slow ema.TFinantial_EMA
	var rsi rsi.TFinantial_RSI
	var ema_vol ema.TFinantial_EMA

	ema_fast.Reset(win_len_min)
	ema_slow.Reset(win_len_max)
	ema_vol.Reset(10)
	rsi.Reset(rsi_win_len, rsi_buy_level, rsi_sell_level)

	var market generic.TMarket

	market.Reset(buycoin, sellcoin, invest)

	var fast_gt_slow = int32(UNDEF)

	iter := 0
	pair := buycoin + sellcoin // btceur

	fmt.Println(pair)

	for {
		time.Sleep(period * time.Second)

		price, err := bitstamp.DoGet(pair)
		if err != nil {
			fmt.Println("Error en el doget de bitstamp")
			continue
		}

		ema_fast.NewPrice(price)
		ema_slow.NewPrice(price)
		fmt.Println("price: ", price, "ema_fast: ", ema_fast.Ema(), "ema_slow: ", ema_slow.Ema(), "time: ", time.Now())

		if iter < training_iters {
			iter++
			continue
		}

		// End of training, start trading

		// Initialize fast_gt_slow only once after training
		if fast_gt_slow == UNDEF {
			if ema_fast.Ema() > ema_slow.Ema() {
				fast_gt_slow = TRUE
			} else {
				fast_gt_slow = FALSE
			}
			fmt.Println("Training ready. Starting trade now...")
			continue
		}

		// fast_gt_slow already defined

		/*
		   if (market.InsideMarket()) {
		       if (price < market.LastBuyPrice()) {
		           market.DoSell(price)
		           fmt.Println("********************************** Activated: CONTROL1")
		           fmt.Println("********************************** VENDE a: ", market.LastSellPrice())
		           fmt.Println("********************************** FIAT: ", market.Fiat())
		       } else {
		           fmt.Println("===> He comprado y esta subiendo, GOOD SIGNAL")
		       }
		   } */

		if fast_gt_slow == FALSE {
			if ema_fast.Ema() < ema_slow.Ema() {
				fmt.Println("ema_fast < ema_slow... Se mantiene la tendencia de bajada")
				// tendency is maintained (falling price)
				continue
			} else {
				if market.InsideMarket() == false {
					if rsi.Buy() {
						market.DoBuy(price)
						fmt.Println("********************************** Buy at: ", market.LastBuyPrice())
						fmt.Println("********************************** CRYPTO: ", market.Crypto())
					} else {
						fmt.Println("Improper RSI to buy: ", rsi.RSI())
					}
				} else {
					fmt.Println("===> Tocaba comprar pero ya estoy dentro")
				}
				fast_gt_slow = TRUE
			}
		} else {
			if ema_fast.Ema() > ema_slow.Ema() {
				fmt.Println("ema_fast > ema_slow... Se mantiene la tendencia de subida")
				// tendency is maintained (climbing price)
				continue
			} else {
				if market.InsideMarket() == true {
					if rsi.Sell() {
						market.DoSell(price)
						fmt.Println("********************************** Sell at: ", market.LastSellPrice())
						fmt.Println("********************************** FIAT: ", market.Fiat())
					} else {
						fmt.Println("Improper RSI to sell: ", rsi.RSI())
					}

				} else {
					fmt.Println("===> Tocaba vender pero estoy fuera")

				}
				fast_gt_slow = FALSE
			}
		}
	}
}