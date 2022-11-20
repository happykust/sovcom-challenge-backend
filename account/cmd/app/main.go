package main

import (
	"account/internal/domain/auth"
	"account/internal/router"
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
	//a := auth.SingUp(account.AccountSignUpRequest{Username: "lofaSdax", Password: "lox", Email: "Mrstarfox29@yandex.ru", FirstName: "dfg", LastName: "dfg", ReferralCode: "90"})
	//fmt.Println(a)
	//gf := auth.SingIn("Mrstarfox29@yandex.ru", "lox")
	//fmt.Println(gf)
	b := auth.VerifyUserRequest(account.AccountVerifyRequest{Id: 11, AdditionalContact: "23423234"})
	fmt.Println(b)
	c := auth.GetVerifyUserStatus(11)
	fmt.Println(c)
	auth.VerifyUser(11, auth.RegistrationStatusVerified)
	//t := auth.SingIn("Mrstarfox29@yandex.ru", "lox")
	//fmt.Println(t)
	go router.AmqpMainRouter()
	select {}

	//select {}
}
