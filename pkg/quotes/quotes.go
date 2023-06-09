package quotes

import (
	"encoding/json"
	"fmt"
	"os"
)

type QuoteJson struct {
	Quotes []Quote `json:"quotes"`
}

type Quote struct {
	Quote string `json:"quote"`
	Author string `json:"author"`
}

func GetQuotes(quotes *[]Quote) error {
	file, err := os.Open("./data/quotes.json")
	if err != nil {
		fmt.Println("Error opening JSON file", err)
		return err
	}
	defer file.Close()

	var quoteJson QuoteJson

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&quoteJson)

	if err != nil {
		fmt.Println("Error decoding JSON file", err)
		return err
	}

	*quotes = quoteJson.Quotes

	return nil
}