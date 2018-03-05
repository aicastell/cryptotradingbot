package ema

import "fmt"

// Finantial indicator: EMA

type TFinantial_EMA struct {
    win_len     int
    k           float64
    ema_ant     float64
    ema_cur     float64
}

func (ob *TFinantial_EMA) Log () {
    fmt.Printf("win_len = %d\n", ob.win_len)
    fmt.Printf("k = %f\n", ob.k)
    fmt.Printf("ema_ant = %f\n", ob.ema_ant)
    fmt.Printf("ema_cur = %f\n", ob.ema_cur)
    fmt.Println("")
}

func (ob *TFinantial_EMA) Reset (win_len int) {
    ob.win_len = win_len
    ob.k = (2.0 / (float64(win_len) + 1.0))
    ob.ema_ant = 0.0
    ob.ema_cur = 0.0
}

func (ob *TFinantial_EMA) WindowLen () int {
    return ob.win_len
}

func (ob *TFinantial_EMA) SetWindowLen (win_len int) {
    ob.win_len = win_len
}

func (ob *TFinantial_EMA) Ema () float64 {
    return ob.ema_cur
}

func (ob *TFinantial_EMA) New_price (price_cur float64) float64 {
    ob.ema_cur = (ob.ema_ant + (ob.k * (price_cur - ob.ema_ant)))
    ob.ema_ant = ob.ema_cur
    return ob.ema_cur
}

