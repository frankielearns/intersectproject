package main

import (
	"fmt"
  "github.com/piquette/finance-go/quote"
)

func main() {
	// 15-min delayed full quote for Apple.
	fmt.Println(quote.Get("XAW.TO"))
}
