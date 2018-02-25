package bitstamp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// {"high": "2559.98", "last": "2559.25", "timestamp": "1497016409", "bid": "2555.00", "vwap": "2492.13", "volume": "1650.07586432", "low": "2423.51", "ask": "2559.25", "open": "2519.00"}

type TBitstampTicker struct {
	High      string `json:"high"`
	Last      string `json:"last"`
	Timestamp string `json:"timestamp"`
	Bid       string `json:"bid"`
	Vwap      string `json:"vwap"`
	Volume    string `json:"volume"`
	Low       string `json:"low"`
	Ask       string `json:"ask"`
	Open      string `json:"open"`
}

// coinpair = btceur
func DoGet(coinpair string) float64 {
	url := fmt.Sprintf("https://www.bitstamp.net/api/v2/ticker/%s/", coinpair)
	fmt.Println(url)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return 0.0
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP res ponse
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return 0.0
	}

	// Print response body
	// bodybytes, _ := ioutil.ReadAll(resp.Body)
	//bodystring := string(bodybytes)
	//	fmt.Println(bodystring)
	//fmt.Println("=======================================================")

	// Fill the record with the data from the JSON
	var bst TBitstampTicker

	// var dat map[string]interface{}
	// json.Unmarshal(bodybytes, &dat)

	// Use json.Decode for reading streams of JSON data
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&bst)
	if err != nil {
		log.Println(err)
		fmt.Println("===========")
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	last, err := strconv.ParseFloat(bst.Last, 64)
	if err != nil {
		log.Fatal("ParseFloat: ", err)
		return 0.0
	}

	return last
}
