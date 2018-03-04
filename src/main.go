package main

import (
    "trading/strategy01"
)

func main() {
    // pair: "btceur"
    // period: 60
    // training_iters 34
    // win_len_min = 13
    // win_len_max = 34
    strategy01.Start("btceur", 60, 34, 13, 34)
}
