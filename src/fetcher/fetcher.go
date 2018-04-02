package fetcher

import (
	"config"
)

type TFetcher interface {
	Create()
	SetConfig(cfg *config.TBotConfig)
	GetPrice(coinpair string) (price float64, err error)
	Destroy()
}
