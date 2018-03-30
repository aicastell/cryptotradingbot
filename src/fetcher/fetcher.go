package fetcher

import (
    "config"
)

type TFetcher interface {
    SetConfig(cfg * config.TBotConfig)
    GetPrice(coinpair string) (price float64, err error)
}

