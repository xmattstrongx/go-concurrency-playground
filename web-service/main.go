package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)

	start := time.Now()

	stockSymbols := []string{
		"googl",
		"msft",
		"bbry",
		"hpq",
		"t",
		"tmus",
		"s",
	}

	var numComplete int

	for _, symbol := range stockSymbols {
		go func(symbol string) {
			resp, _ := http.Get(fmt.Sprintf("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=%s", symbol))
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			quote := new(QuoteResponse)
			xml.Unmarshal(body, &quote)

			fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
			numComplete++
		}(symbol)
	}

	for numComplete < len(stockSymbols) {
		time.Sleep(10 * time.Millisecond)
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

type QuoteResponse struct {
	Status           string
	Name             string
	Symbol           string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	Timestamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChangeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}
