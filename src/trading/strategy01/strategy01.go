package strategy01

import(
	"fmt"
    "time"
	"finantial/ema"
    "markets/generic"
    "markets/bitstamp"
)

var UNDEF = int32(-1)
var TRUE  = int32(1)
var FALSE = int32(0)

func Start(pair string, period time.Duration, training_iters int, win_len_min int, win_len_max int) {

    var ema_low ema.TFinantial_EMA
    var ema_hig ema.TFinantial_EMA
    var market generic.TMarket

    ema_low.Reset(win_len_min)
    ema_hig.Reset(win_len_max)
    market.Reset("bitcoin", "euro", 1000)

    var low_gt_hig = int32(UNDEF)
    var price float64

	iter := 0

    for {
        price = bitstamp.DoGet(pair)

        if price == 0.0 {
            continue
        }

        ema_low.New_price(price)
        ema_hig.New_price(price)
        fmt.Println("price: ", price, "ema_low: ", ema_low.Ema(), "ema_hig: ", ema_hig.Ema(), "time: ", time.Now())

        time.Sleep(period * time.Second)

		if (iter < training_iters) {
		    iter++;
			continue
		}

		// End of training, start trading

        // Initialize low_gt_hig only once after training
        if (low_gt_hig  == UNDEF) {
            if (ema_low.Ema() > ema_hig.Ema()) {
                low_gt_hig = TRUE
            } else {
                low_gt_hig = FALSE
            }
		    fmt.Println("Training ready. Starting trade now...")
            continue
        }

        // log_gt_hig already defined

        if (market.InsideMarket()) {
            if (price < market.LastBuyPrice()) {
                market.DoSell(price)
                fmt.Println("********************************** Activated: CONTROL1")
                fmt.Println("********************************** VENDE a: ", market.LastSellPrice())
                fmt.Println("********************************** FIAT: ", market.Fiat())
            } else {
                fmt.Println("===> He comprado y esta subiendo, GOOD SIGNAL")
            }
        }

        if (low_gt_hig == FALSE) {
            if (ema_low.Ema() < ema_hig.Ema()) {
                fmt.Println("ema_low < ema_hig... Se mantiene la tendencia de bajada")
                // tendency is maintained (BAJADA)
                continue
            } else {
                if (market.InsideMarket() == false) {
                    market.DoBuy(price)
                    low_gt_hig = TRUE
                    fmt.Println("********************************** COMPRA a: ", market.LastBuyPrice())
                    fmt.Println("********************************** CRYPTO: ", market.Crypto())
                } else {
                    fmt.Println("===> Tocaba comprar pero ya estoy dentro")
                }
            }
        } else {
            if (ema_low.Ema() > ema_hig.Ema()) {
                fmt.Println("ema_low < ema_hig... Se mantienee la tendencia de subida")
                // tendency is maintained (SUBIDA)
                continue
            } else {
                if (market.InsideMarket() == true) {
                    market.DoSell(price)
                    low_gt_hig = FALSE
                    fmt.Println("********************************** VENDE a: ", market.LastSellPrice())
                    fmt.Println("********************************** FIAT: ", market.Fiat())
                } else {
                    fmt.Println("===> Tocaba vender pero estoy fuera")

                }
            }
        }
    }
}
