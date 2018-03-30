package stdinput

import (
	"config"
	"fmt"
)

type TStdInput struct {
	Config *config.TBotConfig
}

func (ob *TStdInput) SetConfig(cfg *config.TBotConfig) {
	ob.Config = cfg
}

func (ob *TStdInput) GetPrice(coinpair string) (price float64, err error) {
	fmt.Println("GetPrice en stdinput")
	return 2.0, nil
}
