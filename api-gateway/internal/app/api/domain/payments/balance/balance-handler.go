package balance

import (
	"api-gateway/internal/app/api/domain/payments/balance/delivery/amqp"
	"api-gateway/internal/app/api/token"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"libs/contracts/payments"
	"net/http"
)

func GetBalance(c *gin.Context) {
	payload := token.ValidateAccToken(c)
	if payload == nil || payload.UserVerified == false || payload.Ban == true {
		c.JSON(500, "Invalid token.")
		return
	}
	jsonObject := GetBalanceValidate(c, payload.Id)
	if jsonObject == nil {
		return
	}
	req := amqp.Get(jsonObject)
	objectReq := payments.GetBalancesResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, objectReq)
	return
}
