package simpleDeals

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	simpleDeals "libs/contracts/deals/simple"
	"log"
)

func validateSimpleDealBuy(c *gin.Context, userId uint) []byte {
	var input simpleDeals.SimpleDealBuyRequest

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, "Invalid request body.")
		return []byte(err.Error())
	}

	if input.TickerGroup == "" {
		c.JSON(400, "Invalid ticker group.")
		return []byte("Invalid ticker group.")
	}

	if input.TickerFrom == "" {
		c.JSON(400, "Invalid ticker from.")
		return []byte("Invalid ticker from.")
	}

	if input.TickerTo == "" {
		c.JSON(400, "Invalid ticker to.")
		return []byte("Invalid ticker to.")
	}

	if input.Amount <= 0 {
		c.JSON(400, "Invalid amount.")
		return []byte("Invalid amount.")
	}

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		log.Fatal(errPars)
	}

	return jsonObject
}
