package main

import (
	"example/server/pkg/quotes"
	"fmt"
)

func main() {
	// Get Quotes
	var quotesArray []quotes.Quote
	quotes.GetQuotes(&quotesArray)
	fmt.Println("Quotes", quotesArray)	
}