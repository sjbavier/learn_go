package main

import (
	"cryptomasters/api"
	"fmt"
	"sync"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup // sync loop, wait group for go routines to finish
	for _, currency := range currencies {
		wg.Add(1)                      // open a counter to keep track of open goroutines
		go func(currencyCode string) { // lambda with wg in scope that instantiated a go routine, pattern
			// we use currencyCode param as a COPY to avoid a closure using the same currency in the for loop
			// otherwise we get the last value in the loop for all
			getCurrencyData(currencyCode)
			wg.Done() // remove a counter when a routine is done
		}(currency) // executed immediately, passing in currency from for loop as argument
	}
	wg.Wait() // wait till counters are done
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err != nil {
		println("error")
	}
	fmt.Printf("the rate for %v is %.2f \n", rate.Currency, rate.Price)
}
