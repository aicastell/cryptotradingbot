package rsi

import "fmt"
import "math"

// Finantial indicator: RSI

type TFinantial_RSI struct {
    idx         int
    win_len     int
    price_ant   float64
    price_cur   float64
    variation   float64
    pl          float64
    pg          float64
    rs          float64
    rsi         float64
}

func (ob *TFinantial_RSI) Log () {
    fmt.Printf("win_len = %d\n", ob.win_len)
}

func (ob *TFinantial_RSI) Reset (win_len int) {
    ob.win_len = win_len
}

func (ob *TFinantial_RSI) RSI () float64 {
    return ob.rsi
}

func (ob *TFinantial_RSI) NewPrice (price_cur float64) {
    if (ob.price_ant == 0) && (ob.price_cur == 0) {
        ob.price_cur = price_cur
        ob.idx = 0
        return
    }

    ob.idx++

    ob.price_ant = ob.price_cur
    ob.price_cur = price_cur
    ob.variation = (ob.price_cur - ob.price_ant)

    // Preparativos del primer RSI
    if ob.idx <= ob.win_len {
        if ob.variation > 0 {
            ob.pl = ob.pl + ob.variation
        } else {
            ob.pg = ob.pg + math.Abs(ob.variation)
        }
    }

    // Calculo del primer RSI
    if ob.idx == ob.win_len {
        ob.pl = ob.pl / float64(ob.win_len)
        ob.pg = ob.pg / float64(ob.win_len)
        ob.rs = (ob.pl / ob.pg)
        ob.rsi = (100 - (100 / (1 + ob.rs)))
        return
    }

    // Calculo del resto de RSI's
    if ob.idx > ob.win_len {
        win := 0.0
        lost := 0.0
        // TODO calculos relativos al anterior
        if ob.variation > 0 {
            win = ob.variation
        } else {
            lost = math.Abs(ob.variation)
        }
        ob.pl = (((ob.pl * (float64(ob.win_len)-1.0)) + win) / float64(ob.win_len))
        ob.pg = (((ob.pg * (float64(ob.win_len)-1.0)) + lost) / float64(ob.win_len))
        ob.rs = (ob.pl / ob.pg)
        ob.rsi = (100 - (100 / (1 + ob.rs)))
    }

    return
}

/*
func main() {
    var rsi TFinantial_RSI

    rsi.Reset(14)

    rsi.NewPrice(5.24)
    rsi.NewPrice(5.44)
    rsi.NewPrice(5.42)
    rsi.NewPrice(5.44)
    rsi.NewPrice(5.43)
    rsi.NewPrice(5.45)
    rsi.NewPrice(5.50)
    rsi.NewPrice(5.57)
    rsi.NewPrice(5.66)
    rsi.NewPrice(5.69)
    rsi.NewPrice(5.63)
    rsi.NewPrice(5.63)
    rsi.NewPrice(5.64)
    rsi.NewPrice(5.64)
    rsi.NewPrice(5.65)
    fmt.Println("RSI: ", rsi.RSI())
    rsi.NewPrice(5.63)
    fmt.Println("RSI: ", rsi.RSI())
    rsi.NewPrice(5.63)
    fmt.Println("RSI: ", rsi.RSI())
    rsi.NewPrice(5.68)
    fmt.Println("RSI: ", rsi.RSI())
    rsi.NewPrice(5.72)
    fmt.Println("RSI: ", rsi.RSI())
    rsi.NewPrice(5.70)
    fmt.Println("RSI: ", rsi.RSI())
    rsi.NewPrice(5.68)
    fmt.Println("RSI: ", rsi.RSI())
    rsi.NewPrice(5.60)
    fmt.Println("RSI: ", rsi.RSI())
    rsi.NewPrice(5.41)
    fmt.Println("RSI: ", rsi.RSI())
    rsi.NewPrice(5.28)
    fmt.Println("RSI: ", rsi.RSI())
    rsi.NewPrice(4.98)
    fmt.Println("RSI: ", rsi.RSI())
    rsi.NewPrice(5.13)
    fmt.Println("RSI: ", rsi.RSI())
    rsi.NewPrice(5.10)
    fmt.Println("RSI: ", rsi.RSI())
    rsi.NewPrice(5.07)
    fmt.Println("RSI: ", rsi.RSI())
    rsi.NewPrice(5.25)
    fmt.Println("RSI: ", rsi.RSI())
    rsi.NewPrice(5.17)
    fmt.Println("RSI: ", rsi.RSI())
}
*/
