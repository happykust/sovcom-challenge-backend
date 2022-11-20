package currencyDeals

import (
	"api-gateway/internal/app/api/domain/deals/currencyDeals/delivery/amqp"
	"api-gateway/internal/app/api/token"
	"encoding/json"
	"github.com/gin-gonic/gin"
	currencyDeals "libs/contracts/deals/currency"
	"net/http"
)

func CreateCurrencyDeal(c *gin.Context) {
	payload := token.ValidateAccToken(c)
	if payload == nil || payload.UserVerified == false || payload.Ban == true {
		c.JSON(500, "Invalid token.")
		return
	}

	jsonObject := validateCurrencyDealBuy(c, payload.Id)
	if jsonObject == nil {
		return
	}
	req := amqp.Create(jsonObject)
	objectReq := currencyDeals.CurrencyDealBuyResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, objectReq)
	return
}

func DeleteCurrencyDeal(c *gin.Context) {
	payload := token.ValidateAccToken(c)
	if payload == nil || payload.UserVerified == false || payload.Ban == true {
		c.JSON(500, "Invalid token.")
		return
	}

	jsonObject := validateCurrencyDealDelete(c)
	if jsonObject == nil {
		return
	}
	req := amqp.Delete(jsonObject)
	objectReq := currencyDeals.CurrencyDealSellResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, objectReq)
	return
}

func ReadCurrencyDeal(c *gin.Context) {
	payload := token.ValidateAccToken(c)
	if payload == nil || payload.UserVerified == false || payload.Ban == true {
		c.JSON(500, "Invalid token.")
		return
	}

	jsonObject := validateCurrencyDealRead(payload.Id)
	if jsonObject == nil {
		return
	}
	req := amqp.Read(jsonObject)
	objectReq := currencyDeals.CurrencyDealReadResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, objectReq)
	return
}
