package main

import (
	"account/internal/domain/auth"
	"account/pkg/core/config"
	"account/pkg/core/database"
	"account/pkg/core/database/migrations"
	"fmt"
	"libs/contracts/account"
)

func init() {
	config.Init()
}

func main() {
	database.OpenPostgresConnect()
	migrations.InitMigrations()
	a := auth.SingUp(account.AccountSignUpRequest{Username: "lofdx", Password: "lox", Email: "dfgfdfgfdgd", FirstName: "dfg", LastName: "dfg", ReferralCode: "90"})
	fmt.Println(a)
	gf := auth.SingIn("dfgfdfgfdgd", "lox")
	fmt.Println(gf)
	b := auth.VerifyUserRequest(account.AccountVerifyRequest{Id: 3, AdditionalContact: "23423234"})
	fmt.Println(b)
	c := auth.GetVerifyUserStatus(3)
	fmt.Println(c)
	auth.VerifyUser(3, auth.RegistrationStatusVerified)
	gf = auth.SingIn("dfgfdfgfdgd", "lox")
	fmt.Println(gf)

	select {}
}
