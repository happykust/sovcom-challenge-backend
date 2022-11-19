package auth

import (
	"account/internal/domain/user"
	logger "account/pkg/logging"
	LoggerTypes "account/pkg/logging/types"
	"libs/contracts/account"
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
	// send email
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
	// send email
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
		DeleteUnverifiedUserProfile(userId)
	}
}

func CreatedVerifiedUserAccount(payload user.User) []user.User {
	CreateUserAccount(payload)
	NewUser := FindUserById(payload.ID)
	// send email
	// create balance
	// create RUB wallet ?
	return NewUser

}

func SingIn(email string, password string) account.AccountSignInResponse {
	checkVerifiedUser := FindUserByEmail(email)
	if len(checkVerifiedUser) == 0 {
		oldUnverifiedUser := GetUnverifiedUserByEmail(email)
		if len(oldUnverifiedUser) == 0 {
			return account.AccountSignInResponse{Message: "User not found"}
		}
		validateUserPassword := validatePassword(password, oldUnverifiedUser[0].PasswordHash)
		if !validateUserPassword {
			return account.AccountSignInResponse{Message: "User not found"}
		}
		// send email
		AccessToken, RefreshToken := GenerateTokens(oldUnverifiedUser[0].ID)
		UpdateRefreshTokenUnverifiedUser(oldUnverifiedUser[0].ID, RefreshToken)
		return account.AccountSignInResponse{Message: "Login success", AccessToken: AccessToken, RefreshToken: RefreshToken}
	}
	checkUnVerifiedUser := GetUnverifiedUserByEmail(email)
	if len(checkUnVerifiedUser) != 0 {
		return account.AccountSignInResponse{Message: "User not found"}
	}
	validateUserPassword := validatePassword(password, checkVerifiedUser[0].PasswordHash)
	if !validateUserPassword {
		return account.AccountSignInResponse{Message: "User not found"}
	}
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
