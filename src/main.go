package main

import (
    "trading/strategy01"
)

func main() {
    // buycoin: "btc"
    // sellcoin: "eur"
    // invest: 1000 (eur)
    // period: 60
    // training_iters: 90 minutes
    // win_len_min: 13
    // win_len_max: 34
    strategy01.Start("btc", "eur", 1000, 60, 90, 13, 34)
}
