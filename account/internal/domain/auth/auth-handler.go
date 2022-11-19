package auth

import (
	"account/internal/domain/auth/sendler"
	"account/internal/domain/user"
	logger "account/pkg/logging"
	LoggerTypes "account/pkg/logging/types"
	"encoding/json"
	"fmt"
	"libs/contracts/account"
	"libs/contracts/email"
	"libs/contracts/payments"
)

func SingUp(payload account.AccountSignUpRequest) account.AccountSignUpResponse {
	oldUnverifiedUserEmail := GetUnverifiedUserByEmail(payload.Email)
	oldUnverifiedUserUserName := GetUnverifiedUserByUsername(payload.Username)
	oldUserEmail := FindUserByEmail(payload.Email)
	oldUserUserName := FindUserByUsername(payload.Username)
	if len(oldUnverifiedUserEmail) != 0 && len(oldUnverifiedUserUserName) != 0 {
		return account.AccountSignUpResponse{Message: "User already exists"}
	}
	if len(oldUserEmail) != 0 && len(oldUserUserName) != 0 {
		return account.AccountSignUpResponse{Message: "User already exists"}
	}
	hashPassword := HashUserPassword(payload.Password)
	createdUnverifiedUser := CreatingUnverifiedUser(UnverifiedUsers{Email: payload.Email, PasswordHash: hashPassword, UserName: payload.Username, FirstName: payload.FirstName, LastName: payload.LastName})
	AccessToken, RefreshToken := GenerateTokens(createdUnverifiedUser.ID)
	UpdateRefreshTokenUnverifiedUser(createdUnverifiedUser.ID, RefreshToken)
	emailMessage := email.Request{Email: payload.Email, Subject: "Вы создали акк", Body: "первое сообщени 28"}
	// emailMessage to []byte
	jsonObj, err := json.Marshal(emailMessage)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Marshal error", err)
	}
	sendler.SendEmail(jsonObj)
	return account.AccountSignUpResponse{Message: "User created successfully", AccessToken: AccessToken, RefreshToken: RefreshToken}
}

func VerifyUserRequest(payload account.AccountVerifyRequest) account.AccountVerifyResponse {
	user := GetUnverifiedUserById(payload.Id)
	if len(user) == 0 {
		logger.Log(LoggerTypes.CRITICAL, "User verification failed", nil)
		return account.AccountVerifyResponse{Message: "User verification failed"}
	}
	// get all admins
	// TODO: FIX
	SetAssistants(user[0].ID, 1)
	user = GetUnverifiedUserById(payload.Id)
	// get admin by id
	message := "Вам назначен персональный помощник"
	emailMessage := email.Request{Email: user[0].Email, Subject: "Вы создали акк", Body: "первое сообщени 28"}
	// emailMessage to []byte
	jsonObj, err := json.Marshal(emailMessage)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Marshal error", err)
	}
	sendler.SendEmail(jsonObj)
	// add to admin
	// send email
	return account.AccountVerifyResponse{Message: message, MeetingInformation: message, PersonalAssistant: "olges"}
}

func GetVerifyUserStatus(userId uint) RegistrationStatus {
	userStatus := GetUserVerifiedStatus(userId)
	return userStatus
}

func VerifyUser(userId uint, status RegistrationStatus) {
	localUser := GetUnverifiedUserById(userId)
	if len(localUser) == 0 {
		logger.Log(LoggerTypes.CRITICAL, "User verification failed", nil)
		return
	}
	UpdateUnverifiedUserRegStatus(userId, status)

	localUser = GetUnverifiedUserById(userId)
	if len(localUser) == 0 {
		logger.Log(LoggerTypes.CRITICAL, "User verification failed", nil)
		return
	}

	if localUser[0].RegistrationStatus == RegistrationStatusVerified {
		newUser := user.User{UserName: localUser[0].UserName, Email: localUser[0].Email, FirstName: localUser[0].FirstName, LastName: localUser[0].LastName, PasswordHash: localUser[0].PasswordHash, RefreshTokenHash: localUser[0].RefreshTokenHash}
		CreatedVerifiedUserAccount(newUser)

	}
}

func CreatedVerifiedUserAccount(payload user.User) []user.User {
	CreateUserAccount(payload)
	findNotVerifiedUser := GetUnverifiedUserByEmail(payload.Email)
	DeleteUnverifiedUserProfile(findNotVerifiedUser[0].ID)
	NewUser := FindUserById(payload.ID)
	emailMessage := email.Request{Email: payload.Email, Subject: "Вы создали акк", Body: "первое сообщени 28"}
	// emailMessage to []byte
	jsonObj, err := json.Marshal(emailMessage)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Marshal error", err)
	}
	sendler.SendEmail(jsonObj)
	findNewUser := FindUserByEmail(payload.Email)
	messageToBalance := payments.CreateBalanceUserRequest{UserID: findNewUser[0].ID}
	jsonObj, err = json.Marshal(messageToBalance)
	balance := sendler.SendPayments(jsonObj)
	obj := payments.CreateBalanceUserResponse{}
	err = json.Unmarshal(balance, &obj)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Unmarshal error", err)
	}

	fmt.Println(balance)
	UpdateUserBalance(findNewUser[0].ID, obj.BalanceId)

	// create balance
	// create RUB wallet ?
	return NewUser

}

func SingIn(UserEmail string, password string) account.AccountSignInResponse {
	checkVerifiedUser := FindUserByEmail(UserEmail)
	if len(checkVerifiedUser) == 0 {
		oldUnverifiedUser := GetUnverifiedUserByEmail(UserEmail)
		if len(oldUnverifiedUser) == 0 {
			return account.AccountSignInResponse{Message: "User not found"}
		}
		validateUserPassword := validatePassword(password, oldUnverifiedUser[0].PasswordHash)
		if !validateUserPassword {
			return account.AccountSignInResponse{Message: "User not found"}
		}
		emailMessage := email.Request{Email: UserEmail, Subject: "Вы создали акк", Body: "первое сообщени 28"}
		// emailMessage to []byte
		jsonObj, err := json.Marshal(emailMessage)
		if err != nil {
			logger.Log(LoggerTypes.CRITICAL, "Marshal error", err)
		}
		sendler.SendEmail(jsonObj)
		AccessToken, RefreshToken := GenerateTokens(oldUnverifiedUser[0].ID)
		UpdateRefreshTokenUnverifiedUser(oldUnverifiedUser[0].ID, RefreshToken)
		return account.AccountSignInResponse{Message: "Login success", AccessToken: AccessToken, RefreshToken: RefreshToken}
	}
	checkUnVerifiedUser := GetUnverifiedUserByEmail(UserEmail)
	if len(checkUnVerifiedUser) != 0 {
		return account.AccountSignInResponse{Message: "User not found"}
	}
	validateUserPassword := validatePassword(password, checkVerifiedUser[0].PasswordHash)
	if !validateUserPassword {
		return account.AccountSignInResponse{Message: "User not found"}
	}
	emailMessage := email.Request{Email: UserEmail, Subject: "Вы создали акк", Body: "первое сообщени 28"}
	// emailMessage to []byte
	jsonObj, err := json.Marshal(emailMessage)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Marshal error", err)
	}
	sendler.SendEmail(jsonObj)
	AccessToken, RefreshToken := GenerateTokens(checkVerifiedUser[0].ID)
	UpdateRfToken(checkVerifiedUser[0].ID, RefreshToken)
	return account.AccountSignInResponse{Message: "Login success", AccessToken: AccessToken, RefreshToken: RefreshToken}
}

func Refresh(rfToken string) (string, string) {
	id, err := ParseToken(rfToken)
	if err != nil {
		return "", string(err.Error())
	}
	validateToken := validateRfToken(uint(id), rfToken)
	if !validateToken {
		return "", "Token not valid"
	}
	return RefreshToken(uint(id), rfToken)
}

func SingOut(id uint) account.AccountLogoutResponse {
	DeleteRfToken(id)
	return account.AccountLogoutResponse{Message: "Logout success"}
}
