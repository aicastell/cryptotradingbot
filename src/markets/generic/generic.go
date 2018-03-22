package generic

type TMarket struct {
	coincrypto    string  // bitcoin
	coinfiat      string  // euro
	vcrypto       float64 // 0.0
	vfiat         float64 // 100
	buy           float64 // 0.0
	sell          float64 // 0.0
	fee           float64 // 0.25%
	inside_market bool    // false
}

// bitcoin, euro, 100
func (ob *TMarket) Reset(coincrypto string, coinfiat string, vfiat float64, fee float64) {
	ob.coincrypto = coincrypto
	ob.coinfiat = coinfiat
	ob.vcrypto = 0.0
	ob.vfiat = vfiat
	ob.fee = fee
	ob.buy = 0.0
	ob.inside_market = false
}

func (ob *TMarket) DoBuy(price float64) {
	ob.buy = price
	ob.vcrypto = (ob.vfiat / price) * (1 - (ob.fee / 100))
	ob.vfiat = 0.0
	ob.inside_market = true
}

func (ob *TMarket) DoSell(price float64) {
	ob.sell = price
	ob.vfiat = (ob.vcrypto * price) * (1 - (ob.fee / 100))
	ob.vcrypto = 0.0
	ob.inside_market = false
}

func (ob *TMarket) Fiat() float64 {
	return ob.vfiat
}

func (ob *TMarket) Crypto() float64 {
	return ob.vcrypto
}

func (ob *TMarket) LastBuyPrice() float64 {
	return ob.buy
}

func (ob *TMarket) LastSellPrice() float64 {
	return ob.sell
}

func (ob *TMarket) InsideMarket() bool {
	return ob.inside_market
}
