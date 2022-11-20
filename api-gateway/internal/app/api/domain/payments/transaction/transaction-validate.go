package transaction

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"libs/contracts/payments"
	"log"
	"net/http"
)

func validateCreateTransactionRequest(c *gin.Context, userId uint) []byte {
	var input payments.CreateTransactionRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return []byte(err.Error())
	}

	input.UserId = userId

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		log.Fatal(errPars)
	}

	return jsonObject
}

func validateGetTransactionRequest(c *gin.Context) []byte {
	var input payments.PaymentsTransactionsGet

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return []byte(err.Error())
	}

	if input.TransactionUUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction uuid."})
		return []byte("Invalid uuid.")
	}

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		log.Fatal(errPars)
	}

	return jsonObject
}

func validateUpdateTransactionRequest(c *gin.Context) []byte {
	var input payments.PaymentsTransactionsUpdateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return []byte(err.Error())
	}

	if input.TransactionUUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction uuid."})
		return []byte("Invalid uuid.")
	}

	if input.TransactionStatus == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction status."})
		return []byte("Invalid status.")
	}

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		log.Fatal(errPars)
	}

	return jsonObject
}
