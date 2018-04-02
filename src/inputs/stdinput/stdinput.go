package stdinput

import (
	"config"
	"fmt"
	"time"
)

type TStdInput struct {
	Config *config.TBotConfig
}

func (ob *TStdInput) Create() {
	fmt.Println("Initializing Std Input...")
	// do nothing
}

func (ob *TStdInput) SetConfig(cfg *config.TBotConfig) {
	ob.Config = cfg
}

func (ob *TStdInput) GetPrice(coinpair string) (price float64, err error) {
	fmt.Println("GetPrice en stdinput")

	time.Sleep(time.Duration(ob.Config.Global.Period) * time.Second)

	return 2.0, nil
}

func (ob *TStdInput) Destroy() {
	fmt.Println("Destroying Std Input...")
}
