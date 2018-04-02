package datafile

import (
	"bufio"
	"config"
	"fmt"
	"os"
	"strconv"
	"time"
)

type TDataFile struct {
	Config  *config.TBotConfig
	Scanner *bufio.Scanner
}

func (ob *TDataFile) Create() {
	fmt.Println("Initializing Data File...")

	f, err := os.Open("/tmp/data.dat")
	if err != nil {
		fmt.Println("Data file opened: /tmp/data.dat")
	}

	ob.Scanner = bufio.NewScanner(f)
}

func (ob *TDataFile) SetConfig(cfg *config.TBotConfig) {
	ob.Config = cfg
}

func (ob *TDataFile) GetPrice(coinpair string) (price float64, err error) {
	fmt.Println("GetPrice en datafile")

	time.Sleep(time.Duration(ob.Config.Global.Period) * time.Second)

	ob.Scanner.Scan()

	f, err := strconv.ParseFloat(ob.Scanner.Text(), 64)
	if err != nil {
		fmt.Println("Error converting string to float64")
		// return 0.0, &utils.MyError{15}
		os.Exit(1)
	}

	price = f
	return price, nil
}

func (ob *TDataFile) Destroy() {
	fmt.Println("Destroying Data file...")
}
