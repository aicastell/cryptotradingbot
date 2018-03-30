package stdinput

import (
	"config"
)

type TStdInput struct {
	Config *config.TBotConfig
}

func (ob *TStdInput) SetConfig(cfg *config.TBotConfig) {
	ob.Config = cfg
}

func (ob *TStdInput) GetPrice(coinpair string) (price float64, err error) {
	return price, nil
}
