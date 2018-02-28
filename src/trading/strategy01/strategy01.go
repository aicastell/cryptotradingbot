package strategy01

import(
	"fmt"
    "time"
	"finantial/ema"
    "markets/bitstamp"
)

func Start(pair string, period time.Duration, start_iter int, win_len_min int, win_len_max int) {

    var ema_low ema.TFinantial_EMA
    var ema_hig ema.TFinantial_EMA

    ema_low.Reset(win_len_min)
    ema_hig.Reset(win_len_max)

    var price float64
	iter := 0

    for {
        price = bitstamp.DoGet(pair)

        if price != 0.0 {
            ema_low.New_price(price)
            ema_low.Log()
            ema_hig.New_price(price)
            ema_hig.Log()
        }

        time.Sleep(period * time.Second)
		iter++;

		if (iter < start_iter) {
			continue
		}

		fmt.Println("Training done. Start trading...")

		// START TRADDING HERE!!
    }
}
