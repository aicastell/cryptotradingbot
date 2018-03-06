package exchange

import "time"

type Exchange interface {
	FormatCoinPair(buyCoin, sellCoin string) string
	HasInmediateTraining() bool
	DoGet(coinpair string) (lastValue float64, err error)
	DoGetTrainingValues(period time.Duration, coinpair string) (last24hValues []float64, err error)
}
