package api

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	// if using middlewares, add them here

	return router
}

func AddRoutingGroups(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		quotesGroup := v1.Group("/quotes")
		quotesGroup.GET("/all", GetAllQuotesHandler)
		quotesGroup.GET("/random", GetRandomQuoteHandler)
		quotesGroup.POST(("/add"), AddRandomQuoteHandler)
	}
}