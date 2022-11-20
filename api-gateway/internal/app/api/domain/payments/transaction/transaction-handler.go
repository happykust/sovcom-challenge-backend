package transaction

import (
	"api-gateway/internal/app/api/domain/payments/transaction/delivery/amqp"
	"api-gateway/internal/app/api/token"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"libs/contracts/payments"
	"net/http"
)

func CreateTransaction(c *gin.Context) {
	payload := token.ValidateAccToken(c)
	if payload == nil || payload.UserVerified == false || payload.Ban == true {
		c.JSON(500, "Invalid token.")
		return
	}

	jsonObject := validateCreateTransactionRequest(c, payload.Id)
	if jsonObject == nil {
		return
	}
	req := amqp.Create(jsonObject)
	objectReq := payments.CreateTransactionResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, objectReq)
	return
}

func GetTransaction(c *gin.Context) {
	payload := token.ValidateAccToken(c)
	if payload == nil || payload.UserVerified == false || payload.Ban == true {
		c.JSON(500, "Invalid token.")
		return
	}

	jsonObject := validateGetTransactionRequest(c)
	if jsonObject == nil {
		return
	}
	req := amqp.Get(jsonObject)
	objectReq := payments.PaymentsTransactionsGetResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, objectReq)
	return
}

func UpdateTransactionStatus(c *gin.Context) {
	payload := token.ValidateAccToken(c)
	if payload == nil || payload.UserVerified == false || payload.Ban == true {
		c.JSON(500, "Invalid token.")
		return
	}

	jsonObject := validateUpdateTransactionRequest(c)
	if jsonObject == nil {
		return
	}
	req := amqp.Update(jsonObject)
	objectReq := payments.PaymentsTransactionsUpdateResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, objectReq)
	return
}
