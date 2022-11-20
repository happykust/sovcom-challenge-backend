package promocodes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"libs/contracts/loyality/promocodes"
	"log"
	"net/http"
)

func validatePromocodeCreate(c *gin.Context) []byte {
	var input promocodes.CreateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return []byte(err.Error())
	}

	if input.Promocode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid promocode."})
		return []byte("Invalid promocode.")
	}

	if input.Ticker == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticker."})
		return []byte("Invalid ticker.")
	}

	if input.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount."})
		return []byte("Invalid amount.")
	}

	if input.ActivationCountLimit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activation count limit."})
		return []byte("Invalid activation count limit.")
	}

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		log.Fatal(errPars)
	}

	return jsonObject
}

func validatePromocodeDelete(c *gin.Context) []byte {
	var input promocodes.DeleteRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return []byte(err.Error())
	}

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		log.Fatal(errPars)
	}

	return jsonObject
}

func validatePromocodeUse(c *gin.Context) []byte {
	var input promocodes.UseRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return []byte(err.Error())
	}

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		log.Fatal(errPars)
	}

	return jsonObject
}

func validatePromocodeRead(c *gin.Context) []byte {
	var input promocodes.ReadRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return []byte(err.Error())
	}

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		log.Fatal(errPars)
	}

	return jsonObject
}

func validatePromocodeUpdate(c *gin.Context) []byte {
	var input promocodes.UpdateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return []byte(err.Error())
	}

	if input.Promocode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid promocode."})
		return []byte("Invalid promocode.")
	}

	if input.Ticker == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticker."})
		return []byte("Invalid ticker.")
	}

	if input.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount."})
		return []byte("Invalid amount.")
	}

	if input.ActivationCountLimit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activation count limit."})
		return []byte("Invalid activation count limit.")
	}

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		log.Fatal(errPars)
	}

	return jsonObject
}
