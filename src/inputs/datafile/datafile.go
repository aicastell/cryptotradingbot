package datafile

import (
	"config"
	"fmt"
)

type TDataFile struct {
	Config *config.TBotConfig
}

func (ob *TDataFile) SetConfig(cfg *config.TBotConfig) {
	ob.Config = cfg
}

func (ob *TDataFile) GetPrice(coinpair string) (price float64, err error) {
	fmt.Println("GetPrice en datafile")

	// TODO Abrir fichero, recorrerlo y devolver los distintos float64
	return 23.0, nil
}
