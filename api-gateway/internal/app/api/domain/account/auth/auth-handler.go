package auth

import (
	"api-gateway/internal/app/api/domain/account/auth/delivery/amqp"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"libs/contracts/account"
)

func CreateUserEvent(c *gin.Context) {
	jsonObject := validateUserRegister(c)
	if jsonObject == nil {
		return
	}
	req := amqp.Register(jsonObject)
	objectReq := account.AccountSignInResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.SetCookie("token", objectReq.RefreshToken, 3600, "", "localhost", true, true)
	c.SetCookie("access_token", objectReq.AccessToken, 3600, "", "localhost", true, true)
	c.JSON(200, objectReq)
	return
}

func LoginUserEvent(c *gin.Context) {
	jsonObject := validateUserLogin(c)
	if jsonObject == nil {
		return
	}
	req := amqp.Login(jsonObject)
	objectReq := account.AccountSignInResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.SetCookie("token", objectReq.RefreshToken, 360000, "", "localhost", true, true)
	c.SetCookie("access_token", objectReq.AccessToken, 3600, "", "localhost", true, true)
	c.JSON(200, objectReq)
	return
}

func RefreshTokenEvent(c *gin.Context) {
	jsonObject := validateUserRefresh(c)
	fmt.Println(jsonObject)
	if jsonObject == nil {
		return
	}
	req := amqp.Refresh(jsonObject)
	objectReq := account.AccountSignInResponse{}
	err := json.Unmarshal(req, &objectReq)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.SetCookie("token", objectReq.RefreshToken, 3600, "", "localhost", true, true)
	c.SetCookie("access_token", objectReq.AccessToken, 3600, "", "localhost", true, true)

	c.JSON(200, objectReq)
	return
}

func LogoutUserEvent(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "localhost", true, true)
	c.SetCookie("access_token", "", -1, "", "localhost", true, true)
	c.JSON(200, "Logout")
	return
}

//TODO: Fix cycle imports!
//func VerifyUserRequest(c *gin.Context) {
//	payload := middleware.ValidateAccToken(c)
//	if payload == nil || payload.UserVerified == false || payload.Ban == true {
//		c.JSON(500, "Invalid token.")
//		return
//	}
//
//	jsonObject := validateVerifyRequest(c, payload.Id)
//	if jsonObject == nil {
//		return
//	}
//	req := amqp.Verify(jsonObject)
//	objectReq := account.AccountVerifyResponse{}
//	err := json.Unmarshal(req, &objectReq)
//	if err != nil {
//		c.JSON(500, err)
//		return
//	}
//	c.JSON(http.StatusOK, objectReq)
//	return
//}
