package datafile

import (
    "config"
)

type TDataFile struct {
   Config *config.TBotConfig
}

func (ob *TDataFile) SetConfig(cfg *config.TBotConfig) {
    ob.Config = cfg
}

func (ob *TDataFile) GetPrice(coinpair string) (price float64, err error) {
    return price, err
}

