package main

import (
    "fmt"
    "trading/strategy01"
)

func main() {
    // pair: "btceur"
    // period: 60
    // training_iters: 90 minutes
    // win_len_min: 13
    // win_len_max: 34
    fmt.Println("Empiezo")
    strategy01.Start("btceur", 60, 90, 13, 34)
}
