package router

import "account/internal/domain/auth/delivery"

func AmqpMainRouter() {
	go delivery.CurrencyValidateRequest()
	go delivery.SupportValidateRequest()
	go delivery.GetAllVUser()
	go delivery.GetAllNVUser()
	go delivery.Refresh()
	go delivery.SignUp()
	go delivery.SignIn()
	go delivery.VerifyRequest()
	go delivery.Approve()
	select {}
}
