package promocodes

import (
	"api-gateway/internal/app/api/domain/loyality/promocodes/delivery/amqp"
	"api-gateway/internal/app/api/token/middleware"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"libs/contracts/loyality/promocodes"
	"libs/contracts/payments"
	"net/http"
)

func CreatePromocode(c *gin.Context) {
	payload := middleware.ValidateAccToken(c)
	if payload == nil || payload.UserVerified == false || payload.Ban == true || payload.Role == "user" {
		c.JSON(500, "Invalid token.")
		return
	}

	jsonObject := validatePromocodeCreate(c)
	if jsonObject == nil {
		return
	}
	req := amqp.Create(jsonObject)
	objectReq := payments.GetBalancesResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, objectReq)
	return
}

func DeletePromocode(c *gin.Context) {
	payload := middleware.ValidateAccToken(c)
	if payload == nil || payload.UserVerified == false || payload.Ban == true || payload.Role == "user" {
		c.JSON(500, "Invalid token.")
		return
	}

	jsonObject := validatePromocodeDelete(c)
	if jsonObject == nil {
		return
	}
	req := amqp.Delete(jsonObject)
	objectReq := promocodes.DeleteResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, objectReq)
	return
}

func ReadPromocode(c *gin.Context) {
	payload := middleware.ValidateAccToken(c)
	if payload == nil || payload.UserVerified == false || payload.Ban == true || payload.Role == "user" {
		c.JSON(500, "Invalid token.")
		return
	}

	jsonObject := validatePromocodeRead(c)
	if jsonObject == nil {
		return
	}
	req := amqp.Read(jsonObject)
	objectReq := promocodes.ReadResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, objectReq)
	return
}

func UpdatePromocode(c *gin.Context) {
	payload := middleware.ValidateAccToken(c)
	if payload == nil || payload.UserVerified == false || payload.Ban == true || payload.Role == "user" {
		c.JSON(500, "Invalid token.")
		return
	}

	jsonObject := validatePromocodeUpdate(c)
	if jsonObject == nil {
		return
	}
	req := amqp.Update(jsonObject)
	objectReq := promocodes.UpdateResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, objectReq)
	return
}

func UsePromocode(c *gin.Context) {
	payload := middleware.ValidateAccToken(c)
	if payload == nil || payload.UserVerified == false || payload.Ban == true {
		c.JSON(500, "Invalid token.")
		return
	}

	jsonObject := validatePromocodeUse(c)
	if jsonObject == nil {
		return
	}
	req := amqp.Use(jsonObject)
	objectReq := promocodes.UseResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, objectReq)
	return
}
