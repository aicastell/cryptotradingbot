package sma

import "fmt"

// Finantial indicator: SMA

type TFinantial_SMA struct {
	win_len int
	sma     float64
	values  []float64
}

func (ob *TFinantial_SMA) Log() {
	fmt.Println("sma: ", ob.sma)
	fmt.Println("len: ", len(ob.values))
	fmt.Println("win_len: ", ob.win_len)
	fmt.Println("values: ", ob.values)
	fmt.Println("")
}

func (ob *TFinantial_SMA) Reset(win_len int) {
	ob.sma = 0.0
	ob.win_len = win_len
	ob.values = make([]float64, 0, win_len)
}

func (ob *TFinantial_SMA) SMA() float64 {
	return ob.sma
}

func (ob *TFinantial_SMA) NewPrice(price_cur float64) {

	if len(ob.values) < ob.win_len {
		// Fill-in the empty slice values
		ob.values = append(ob.values, price_cur)
		if len(ob.values) < ob.win_len {
			return
		}
	} else {
		// Make the LIFO rotation
		ob.values = ob.values[1:]
		ob.values = append(ob.values, price_cur)
	}

	// Get SMA
	ob.sma = 0.0
	for i := 0; i < ob.win_len; i++ {
		ob.sma = ob.sma + ob.values[i]
	}
	ob.sma = (ob.sma / float64(ob.win_len))
	return
}

/*
func main() {
	var sma TFinantial_SMA

	sma.Reset(10)

    for i:= 0; i < 30; i++ {
    	sma.NewPrice(float64(i))
    	sma.Log()
    }
}
*/
