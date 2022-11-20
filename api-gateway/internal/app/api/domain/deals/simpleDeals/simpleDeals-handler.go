package simpleDeals

import (
	"api-gateway/internal/app/api/domain/deals/simpleDeals/delivery/amqp"
	"api-gateway/internal/app/api/token"
	"encoding/json"
	"github.com/gin-gonic/gin"
	simpleDeals "libs/contracts/deals/simple"
	"net/http"
)

func CreateSimpleDeal(c *gin.Context) {
	payload := token.ValidateAccToken(c)
	if payload == nil || payload.UserVerified == false || payload.Ban == true {
		c.JSON(500, "Invalid token.")
		return
	}

	jsonObject := validateSimpleDealBuy(c, payload.Id)
	if jsonObject == nil {
		return
	}
	req := amqp.Create(jsonObject)
	objectReq := simpleDeals.SimpleDealBuyResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, objectReq)
	return
}
