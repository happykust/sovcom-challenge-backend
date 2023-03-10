package auth

import (
	"account/internal/domain/auth/sendler"
	"account/internal/domain/user"
	logger "account/pkg/logging"
	LoggerTypes "account/pkg/logging/types"
	"encoding/json"
	"fmt"
	"libs/contracts/account"
	"libs/contracts/currency/currencyToAccounts"
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
	AccessToken, RefreshToken := GenerateTokens(createdUnverifiedUser.ID, createdUnverifiedUser.Ban, false, user.RoleUser)
	UpdateRefreshTokenUnverifiedUser(createdUnverifiedUser.ID, RefreshToken)
	emailMessage := email.Request{Email: payload.Email, Subject: "Поздравляем, Ваш аккаунт зарегистрирован", Body: "Для взаимодействия с основными функционалом платформы нужно верифицировать аккаунт, сделать это можно в профиле "}
	jsonObj, err := json.Marshal(emailMessage)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Marshal error", err)
	}
	sendler.SendEmail(jsonObj)
	return account.AccountSignUpResponse{Message: "User created successfully", AccessToken: AccessToken, RefreshToken: RefreshToken}
}

func BanUser(userId uint, BanStatus bool) bool {
	BanUserStatusUpdate(userId, BanStatus)
	userBanStatus := GetUserBanStatus(userId)
	if userBanStatus == true {
		return true
	} else {
		return false
	}
}

func VerifyUserRequest(payload account.AccountVerifyRequest) account.AccountVerifyResponse {
	user := GetUnverifiedUserById(payload.Id)
	if len(user) == 0 {
		logger.Log(LoggerTypes.CRITICAL, "User verification failed", nil)
		return account.AccountVerifyResponse{Message: "User verification failed"}
	}

	SetAssistants(user[0].ID, 1)
	user = GetUnverifiedUserById(payload.Id)
	message := "Вам назначен персональный помощник"
	emailMessage := email.Request{Email: user[0].Email, Subject: "Вы отправили запрос на верификацию!", Body: "Ожидайте, скоро с Вами свяжется администратор"}
	jsonObj, err := json.Marshal(emailMessage)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Marshal error", err)
	}
	sendler.SendEmail(jsonObj)
	return account.AccountVerifyResponse{Message: message, MeetingInformation: message, PersonalAssistant: "olges"}
}

func GetVerifyUserStatus(userId uint) RegistrationStatus {
	userStatus := GetUserVerifiedStatus(userId)
	return userStatus
}

func VerifyUser(userId uint, status RegistrationStatus) bool {
	localUser := GetUnverifiedUserById(userId)
	if len(localUser) == 0 {
		logger.Log(LoggerTypes.CRITICAL, "User verification failed", nil)
		return false
	}
	UpdateUnverifiedUserRegStatus(userId, status)

	localUser = GetUnverifiedUserById(userId)
	if len(localUser) == 0 {
		logger.Log(LoggerTypes.CRITICAL, "User verification failed", nil)
		return false
	}

	if localUser[0].RegistrationStatus == RegistrationStatusVerified {
		newUser := user.User{UserName: localUser[0].UserName, Email: localUser[0].Email, FirstName: localUser[0].FirstName, LastName: localUser[0].LastName, PasswordHash: localUser[0].PasswordHash, RefreshTokenHash: localUser[0].RefreshTokenHash}
		CreatedVerifiedUserAccount(newUser)
		return true
	}
	return false
}

func CreatedVerifiedUserAccount(payload user.User) []user.User {
	CreateUserAccount(payload)
	findNotVerifiedUser := GetUnverifiedUserByEmail(payload.Email)
	DeleteUnverifiedUserProfile(findNotVerifiedUser[0].ID)
	NewUser := FindUserById(payload.ID)
	emailMessage := email.Request{Email: payload.Email, Subject: "Ваш аккаунт верифицирован! ", Body: "Перейдите в профиль для более подробной информации"}
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
		emailMessage := email.Request{Email: UserEmail, Subject: "В Ваш аккаунт выполнен вход", Body: "Это были Вы? Подробнее можно узнать в профиле"}
		jsonObj, err := json.Marshal(emailMessage)
		if err != nil {
			logger.Log(LoggerTypes.CRITICAL, "Marshal error", err)
		}
		sendler.SendEmail(jsonObj)
		AccessToken, RefreshToken := GenerateTokens(oldUnverifiedUser[0].ID, oldUnverifiedUser[0].Ban, false, user.RoleUser)
		UpdateRefreshTokenUnverifiedUser(oldUnverifiedUser[0].ID, RefreshToken)
		fmt.Println("auth NOOOOOOverified")
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
	jsonObj, err := json.Marshal(emailMessage)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Marshal error", err)
	}
	sendler.SendEmail(jsonObj)
	AccessToken, RefreshToken := GenerateTokens(checkVerifiedUser[0].ID, checkVerifiedUser[0].Ban, true, user.RoleUser)
	UpdateRfToken(checkVerifiedUser[0].ID, RefreshToken)
	fmt.Println("auth verified")
	return account.AccountSignInResponse{Message: "Login success", AccessToken: AccessToken, RefreshToken: RefreshToken}
}

func GetAllUnVerifyUsers() []account.AccountGetOneNVUserResponse {
	usersNV := GetAllUnverifiedUsers()
	var users []account.AccountGetOneNVUserResponse
	for _, user := range usersNV {
		users = append(users, account.AccountGetOneNVUserResponse{UserName: user.UserName, Email: user.Email, FirstName: user.FirstName, LastName: user.LastName, RegistrationStatus: string(user.RegistrationStatus), Ban: user.Ban, PersonalAssistant: user.PersonalAssistant, MeetingInformation: user.MeetingInformation, ReferralCode: user.ReferralCode, AdditionalContact: user.AdditionalContact})
	}
	return users
}

func GetAllVerifyUser() []account.AccountGetOneVUserResponse {
	usersV := GetAllUsers()
	var users []account.AccountGetOneVUserResponse
	for _, user := range usersV {
		users = append(users, account.AccountGetOneVUserResponse{UserName: user.UserName, Email: user.Email, FirstName: user.FirstName, LastName: user.LastName, Ban: user.Ban, BalanceId: user.BalanceId, ReferralId: user.ReferralId, Role: string(user.Role), ReferralCode: user.ReferralCode})
	}
	return users
}

func UpdateUserRole(id uint, role user.Role) {
	UpdateUserRoleStatus(id, role)
}

func GetUserRole(id uint) user.Role {
	return GetUserRoleStatus(id)
}

func ValidateTokens(accessToken string) currencyToAccounts.ValidateResponse {
	t, err := ParseToken(accessToken)
	if err != nil || t == nil || t.Ban == true {
		return currencyToAccounts.ValidateResponse{Status: false}
	}
	return currencyToAccounts.ValidateResponse{Status: true, UserID: t.Id}
}

func Refresh(rfToken string) (string, string) {
	claims, err := ParseToken(rfToken)
	if err != nil {
		return "", string(err.Error())
	}
	validateToken := validateRfToken(claims.Id, claims.UserVerified, rfToken)
	if !validateToken {
		return "", "Token not valid"
	}
	return RefreshToken(claims.Id, claims.Ban, claims.Role, claims.UserVerified, rfToken)
}

func SingOut(id uint) account.AccountLogoutResponse {
	DeleteRfToken(id)
	return account.AccountLogoutResponse{Message: "Logout success"}
}
