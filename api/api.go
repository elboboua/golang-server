package api

import (
	"encoding/json"
	"example/server/pkg/quotes"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get Quotes
var quotesArray []quotes.Quote

func getQuotes() []quotes.Quote {
	if len(quotesArray) == 0 {
		 quotes.GetQuotes(&quotesArray)
	}
	return quotesArray
}

func getRandomInt(limit int) int {
	randomNumber := rand.Intn(limit - 1)
	return randomNumber
}

func GetAllQuotesHandler(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK,getQuotes())
}

func GetRandomQuoteHandler(ctx * gin.Context) {
	randomIndex := getRandomInt(len(getQuotes()))
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, quotesArray[randomIndex])
}

func AddRandomQuoteHandler(ctx *gin.Context) {
	var q quotes.Quote
	err := json.NewDecoder(ctx.Request.Body).Decode(&q)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	// ensure all fields are filled out for struct 
	if q.Quote == "" || q.Author == "" {
		ctx.String(http.StatusBadRequest, "Missing quote or author in request");
		return
	}
	// ensure that the quote does not already exist
	for _, quote := range getQuotes() {
		if quote.Quote == q.Quote {
			ctx.String(http.StatusConflict, "quote already exists")
			return
		}
	}

	quotesArray = append(quotesArray, q)
	ctx.Status(http.StatusOK)
}