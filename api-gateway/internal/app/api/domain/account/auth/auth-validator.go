package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"libs/contracts/account"
	"log"
	"net/http"
	"net/mail"
)

func validateUserRegister(c *gin.Context) []byte {
	var input account.AccountSignUpRequest

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return nil
	}

	_, errMail := mail.ParseAddress(input.Email)
	if errMail != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMail.Error()})
		return nil
	}

	if len(input.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 6 characters"})
		return nil
	}

	if len(input.Username) != 7 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username id must be 7 characters"})
		return nil
	}

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		log.Fatal(errPars)
	}

	return jsonObject
}

func validateUserLogin(c *gin.Context) []byte {
	var input account.AccountSignInRequest

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return nil
	}

	_, errMail := mail.ParseAddress(input.Email)
	if errMail != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMail.Error()})
		return nil
	}

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		log.Fatal(errPars)
	}

	return jsonObject
}

func validateUserRefresh(c *gin.Context) []byte {

	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(500, err)
		return nil
	}
	fmt.Println(token)
	var input account.AccountRefreshRequest

	input.RefreshToken = token
	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		fmt.Println("err")
	}

	return jsonObject

}

func validateUserLogout(c *gin.Context) []byte {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(500, err)
		return nil
	}
	fmt.Println(token)
	var input account.AccountRefreshRequest

	input.RefreshToken = token
	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		fmt.Println("err")
	}

	return jsonObject

}

func validateVerifyRequest(c *gin.Context, userId uint) []byte {
	var input account.AccountVerifyRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return []byte(err.Error())
	}

	input.Id = userId

	if input.AdditionalContact == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid additional contact."})
		return []byte("Invalid additional contact.")
	}

	jsonObject, errPars := json.Marshal(input)
	if errPars != nil {
		fmt.Println("err")
	}
	return jsonObject
}
