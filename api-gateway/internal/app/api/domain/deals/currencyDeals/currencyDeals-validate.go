package currencyDeals

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	CurrencyDeals "libs/contracts/deals/currency"
	"log"
	"net/http"
)

const (
	userId       = "userId"
	userBan      = "userBan"
	userVerified = "userVerified"
	userRole     = "userRole"
)

func validateCurrencyDealBuy(c *gin.Context, userId uint) []byte {
	var input CurrencyDeals.CurrencyDealBuyRequest

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return []byte(err.Error())
	}

	if input.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount."})
		return []byte("Invalid amount.")
	}

	if input.Currency == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid currency."})
		return []byte("Invalid currency.")
	}

	input.UserID = userId

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		log.Fatal(errPars)
	}

	return jsonObject
}

func validateCurrencyDealDelete(c *gin.Context) []byte {
	var input CurrencyDeals.CurrencyDealDeleteRequest

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return []byte(err.Error())
	}

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		log.Fatal(errPars)
	}

	return jsonObject
}

func validateCurrencyDealRead(userId uint) []byte {
	var input CurrencyDeals.CurrencyDealReadRequest
	input.UserID = userId

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		log.Fatal(errPars)
	}

	return jsonObject
}
