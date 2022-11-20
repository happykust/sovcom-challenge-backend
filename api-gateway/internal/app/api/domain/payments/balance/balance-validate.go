package balance

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"libs/contracts/payments"
	"log"
	"net/http"
)

func GetBalanceValidate(c *gin.Context, userId uint) []byte {
	var input payments.GetBalancesRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return []byte(err.Error())
	}

	input.UserID = userId

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		log.Fatal(errPars)
	}

	return jsonObject
}
