package strategy01

import (
	"config"
	"fetcher"
	"finantial/ema"
	"fmt"
	"market"
	"time"
)

const (
	UNDEF = int32(-1)
	TRUE  = int32(1)
	FALSE = int32(0)
)

type TStrag01 struct{}

func (ob *TStrag01) Start(cfg *config.TBotConfig, input *fetcher.TFetcher) {

	var ema_fast ema.TFinantial_EMA
	var ema_slow ema.TFinantial_EMA
	var ema_vol ema.TFinantial_EMA

	ema_fast.Reset(cfg.EMA.Fast)
	ema_slow.Reset(cfg.EMA.Slow)
	ema_vol.Reset(10)

	var mm market.TMarket

	mm.Reset(cfg.Global.BuyCoin, cfg.Global.SellCoin, cfg.Global.Invest, cfg.Global.Fee)

	var fast_gt_slow = int32(UNDEF)

	iter := 0
	pair := cfg.Global.BuyCoin + cfg.Global.SellCoin // btceur

	fmt.Println(pair)

	for {
		price, err := (*input).GetPrice(pair)
		if err != nil {
			fmt.Println("Error en el metodo GetPrice")
			continue
		}

		ema_fast.NewPrice(price)
		ema_slow.NewPrice(price)
		fmt.Println("price: ", price, "ema_fast: ", ema_fast.Ema(), "ema_slow: ", ema_slow.Ema(), "time: ", time.Now())

		if iter < cfg.Global.TrainingIters {
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
		   if (mm.InsideMarket()) {
		       if (price < mm.LastBuyPrice()) {
		           mm.DoSell(price)
		           fmt.Println("********************************** Activated: CONTROL1")
		           fmt.Println("********************************** VENDE a: ", mm.LastSellPrice())
		           fmt.Println("********************************** FIAT: ", mm.Fiat())
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
				if mm.InsideMarket() == false {
					mm.DoBuy(price)
					fmt.Println("********************************** Buy at: ", mm.LastBuyPrice())
					fmt.Println("********************************** CRYPTO: ", mm.Crypto())
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
				if mm.InsideMarket() == true {
					mm.DoSell(price)
					fmt.Println("********************************** Sell at: ", mm.LastSellPrice())
					fmt.Println("********************************** FIAT: ", mm.Fiat())
				} else {
					fmt.Println("===> Tocaba vender pero estoy fuera")

				}
				fast_gt_slow = FALSE
			}
		}
	}
}
