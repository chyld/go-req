package main

import (
	"fmt"

	"github.com/chyld/alpha/request"
)

func main() {
	ch := make(chan map[string]interface{})
	symbols := []string{"AAPL", "GOOG", "MSFT", "AMZN"}
	count := 0

	for _, symbol := range symbols {
		go request.Request("http://dev.markitondemand.com/Api/v2/Quote/json?symbol="+symbol, ch)

	}

	for {
		quote := <-ch
		count++
		fmt.Printf("quote: %#v\n\n", quote)

		if count == len(symbols) {
			break
		}
	}
}
