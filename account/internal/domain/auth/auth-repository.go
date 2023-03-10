package auth

import (
	"account/internal/domain/user"
	"account/pkg/core/database"
	"libs/contracts/account"
)

func CreatingUnverifiedUser(userData UnverifiedUsers) UnverifiedUsers {
	database.PG.Create(&userData)
	return userData
}

func GetUnverifiedUserByEmail(email string) []UnverifiedUsers {
	var findUser []UnverifiedUsers
	database.PG.Find(&findUser, "email = ?", email)
	return findUser
}

func GetUnverifiedUserByUsername(userName string) []UnverifiedUsers {
	var findUser []UnverifiedUsers
	database.PG.Find(&findUser, "user_name = ?", userName)
	return findUser
}

func GetUnverifiedUserById(id uint) []UnverifiedUsers {
	var findUser []UnverifiedUsers
	database.PG.Find(&findUser, "id = ?", id)
	return findUser
}

func UpdateUnverifiedUser(userData UnverifiedUsers) UnverifiedUsers {
	database.PG.Model(&userData).Where("id = ?", userData.ID).Updates(userData)
	return userData
}

func CreateUserAccount(userData user.User) {
	database.PG.Create(&userData)
}

func FindUserById(id uint) []user.User {
	var findUser []user.User
	database.PG.Find(&findUser, "id = ?", id)
	return findUser
}

func FindUserByEmail(email string) []user.User {
	var findUser []user.User
	database.PG.Find(&findUser, "email = ?", email)
	return findUser
}

func FindUserByUsername(userName string) []user.User {
	var findUser []user.User
	database.PG.Find(&findUser, "user_name = ?", userName)
	return findUser
}

func GetUserVerifiedStatus(userId uint) RegistrationStatus {
	var status RegistrationStatus
	database.PG.Raw("SELECT registration_status FROM unverified_users WHERE id = ?", userId).Scan(&status)
	return status

}

func UpdateUnverifiedUserRegStatus(userId uint, status RegistrationStatus) {
	database.PG.Exec("UPDATE unverified_users SET registration_status = ? WHERE id = ?", status, userId)
}

func UpdateRefreshTokenUnverifiedUser(id uint, refreshToken string) bool {
	database.PG.Exec("UPDATE unverified_users SET refresh_token_hash = ? WHERE id = ?", refreshToken, id)
	return true
}
func DeleteRfToken(id uint) bool {
	database.PG.Exec("UPDATE users SET refresh_token_hash = '' WHERE id = ?", id)
	return true
}

func DeleteUnverifiedUserProfile(id uint) bool {
	// delete by id
	database.PG.Delete(&UnverifiedUsers{}, "id = ?", id)
	return true
}

func BanUserStatusUpdate(userId uint, ban bool) bool {
	database.PG.Exec("UPDATE users SET ban = ? WHERE id = ?", ban, userId)
	return true
}
func GetAllUnverifiedUsers() []account.AccountGetOneNVUserResponse {
	var usersI []UnverifiedUsers
	database.PG.Find(&usersI)
	var users []account.AccountGetOneNVUserResponse
	for _, user := range usersI {
		users = append(users, account.AccountGetOneNVUserResponse{UserName: user.UserName, Email: user.Email,
			FirstName: user.FirstName, LastName: user.LastName, Ban: user.Ban, ReferralCode: user.ReferralCode})
	}
	return users
}

func GetAllUsers() []user.User {
	var users []user.User
	database.PG.Find(&users)
	return users
}

func UpdateUserRoleStatus(userId uint, role user.Role) bool {
	database.PG.Exec("UPDATE users SET role = ? WHERE id = ?", role, userId)
	return true
}

func GetUserRoleStatus(userId uint) user.Role {
	var status user.Role
	database.PG.Raw("SELECT role FROM users WHERE id = ?", userId).Scan(&status)
	return status
}

func GetUserBanStatus(userId uint) bool {
	var status bool
	database.PG.Raw("SELECT ban FROM users WHERE id = ?", userId).Scan(&status)
	return status

}
func UpdateUserBalance(userId uint, balanceId uint) {
	database.PG.Exec("UPDATE users SET balance_id = ? WHERE id = ?", balanceId, userId)
}

func SetAssistants(userId uint, assistant uint) {
	database.PG.Exec("UPDATE unverified_users SET personal_assistant = ? WHERE id = ?", assistant, userId)
}

func UpdateRfToken(id uint, refreshToken string) bool {
	database.PG.Exec("UPDATE users SET refresh_token_hash = ? WHERE id = ?", refreshToken, id)
	return true
}

func DeleteUnverifiedUser(id uint) bool {
	database.PG.Delete(&UnverifiedUsers{}, "id = ?", id)
	return true
}
