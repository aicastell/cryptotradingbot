package main

import (
	"finantial/ema"
	"markets/bitstamp"
	"time"
)

func main() {
	var ema_low ema.TFinantial_EMA
	var ema_hig ema.TFinantial_EMA

	ema_low.Reset(13)
	ema_hig.Reset(34)

	var euros float64

	for {
		euros = bitstamp.DoGet("btceur")

        if euros != 0.0 {
    		ema_low.New_price(euros)
    		ema_low.Log()
    		ema_hig.New_price(euros)
    		ema_hig.Log()
        }

		time.Sleep(60 * time.Second)
	}
}
