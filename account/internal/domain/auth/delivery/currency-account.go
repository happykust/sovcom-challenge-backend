package delivery

import (
	"account/internal/domain/auth"
	amqp_easier "account/pkg/core/broker/amqp-easier"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"libs/contracts/account"
	"libs/contracts/currency"
	"libs/contracts/currency/currencyToAccounts"
)

func SupportValidateRequest() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(account.SupportValidateConsumerName,
		account.AccountExchange, "topic", account.SupportValidateTopic, account.SupportToAccountsQueueName)
	defer func(amqpChannel *amqp.Channel) {
		err := amqpChannel.Close()
		if err != nil {

		}
	}(amqpChannel)
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	stopChan := make(chan bool)
	go func() {
		for d := range messageChannel {
			validateReq := &currencyToAccounts.ValidateRequest{}
			err := json.Unmarshal(d.Body, validateReq)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(validateReq)
			res := auth.ValidateTokens(validateReq.AccessToken)
			t, err := json.Marshal(&currencyToAccounts.ValidateResponse{Status: res.Status, UserID: res.UserID})
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          t,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
	<-stopChan
}

func CurrencyValidateRequest() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(currency.CurrencyToAccountsConsumerName,
		account.AccountExchange, "topic", currency.CurrencyToAccountsRoutingKey, currency.CurrencyToAccountsQueueName)
	defer func(amqpChannel *amqp.Channel) {
		err := amqpChannel.Close()
		if err != nil {

		}
	}(amqpChannel)
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	stopChan := make(chan bool)
	go func() {
		for d := range messageChannel {
			validateReq := &currencyToAccounts.ValidateRequest{}
			err := json.Unmarshal(d.Body, validateReq)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(validateReq)
			res := auth.ValidateTokens(validateReq.AccessToken)
			t, err := json.Marshal(&currencyToAccounts.ValidateResponse{Status: res.Status, UserID: res.UserID})
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          t,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
	<-stopChan
}

func GetAllVUser() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(account.GetAllVUserConsumerName,
		account.AccountExchange, "topic", account.GetAllVUserTopic, account.GetAllVUserQueueName)
	defer func(amqpChannel *amqp.Channel) {
		err := amqpChannel.Close()
		if err != nil {

		}
	}(amqpChannel)
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	stopChan := make(chan bool)
	go func() {
		for d := range messageChannel {
			fmt.Println(string(d.Body))
			getAllVUser := &account.AccountGetAllVUsersRequest{}
			err := json.Unmarshal(d.Body, getAllVUser)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(getAllVUser)
			users := auth.GetAllVerifyUser()
			t, err := json.Marshal(&account.AccountGetAllVUsersResponse{Users: users})
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          t,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
	<-stopChan
}

func GetAllNVUser() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(account.GetAllNVUserConsumerName,
		account.AccountExchange, "topic", account.GetAllNVUserTopic, account.GetAllNVUserQueueName)
	defer func(amqpChannel *amqp.Channel) {
		err := amqpChannel.Close()
		if err != nil {

		}
	}(amqpChannel)
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	stopChan := make(chan bool)
	go func() {
		for d := range messageChannel {
			fmt.Println(string(d.Body))
			getAllNVUser := &account.AccountGetAllVUsersRequest{}
			err := json.Unmarshal(d.Body, getAllNVUser)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(getAllNVUser)
			users := auth.GetAllUnverifiedUsers()
			t, err := json.Marshal(&account.AccountGetAllNVUsersResponse{Users: users})
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          t,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
	<-stopChan

}

func Refresh() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(account.RefreshConsumerName,
		account.AccountExchange, "topic", account.RefreshTopic, account.RefreshQueueName)
	defer func(amqpChannel *amqp.Channel) {
		err := amqpChannel.Close()
		if err != nil {

		}
	}(amqpChannel)
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	stopChan := make(chan bool)
	go func() {
		for d := range messageChannel {
			fmt.Println(string(d.Body))
			refresh := &account.AccountRefreshRequest{}
			err := json.Unmarshal(d.Body, refresh)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(refresh)
			res, ress := auth.Refresh(refresh.RefreshToken)
			t, err := json.Marshal(&account.AccountRefreshResponse{AccessToken: res, RefreshToken: ress})
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          t,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
	<-stopChan
}

func SignUp() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(account.SignUpConsumerName,
		account.AccountExchange, "topic", account.SignUpTopic, account.SignUpQueueName)
	defer func(amqpChannel *amqp.Channel) {
		err := amqpChannel.Close()
		if err != nil {

		}
	}(amqpChannel)
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	stopChan := make(chan bool)
	go func() {
		for d := range messageChannel {
			fmt.Println(string(d.Body))
			signUp := &account.AccountSignUpRequest{}
			err := json.Unmarshal(d.Body, signUp)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(signUp)
			res := auth.SingUp(*signUp)
			t, err := json.Marshal(&account.AccountSignUpResponse{Message: res.Message, RefreshToken: res.RefreshToken,
				AccessToken: res.AccessToken})
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          t,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
	<-stopChan
}

func SignIn() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(account.SignInConsumerName,
		account.AccountExchange, "topic", account.SignIn, account.SignInQueueName)
	defer func(amqpChannel *amqp.Channel) {
		err := amqpChannel.Close()
		if err != nil {

		}
	}(amqpChannel)
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	stopChan := make(chan bool)
	go func() {
		for d := range messageChannel {
			fmt.Println(string(d.Body))
			signIn := &account.AccountSignInRequest{}
			err := json.Unmarshal(d.Body, signIn)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(signIn)
			res := auth.SingIn(signIn.Email, signIn.Password)
			t, err := json.Marshal(&account.AccountSignInResponse{Message: res.Message, RefreshToken: res.RefreshToken,
				AccessToken: res.AccessToken})
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          t,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
	<-stopChan
}

func VerifyRequest() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(account.VerifyConsumerName,
		account.AccountExchange, "topic", account.VerifyTopic, account.VerifyQueueName)
	defer func(amqpChannel *amqp.Channel) {
		err := amqpChannel.Close()
		if err != nil {

		}
	}(amqpChannel)
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	stopChan := make(chan bool)
	go func() {
		for d := range messageChannel {
			fmt.Println(string(d.Body))
			verify := &account.AccountVerifyRequest{}
			err := json.Unmarshal(d.Body, verify)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(verify)
			res := auth.VerifyUserRequest(*verify)
			t, err := json.Marshal(&account.AccountVerifyResponse{Message: res.Message,
				MeetingInformation: res.MeetingInformation, PersonalAssistant: res.PersonalAssistant})
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          t,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
	<-stopChan
}

func Approve() {
	messageChannel, amqpChannel, conn := amqp_easier.ConsumerConstructor(account.ApproveConsumerName,
		account.AccountExchange, "topic", account.ApproveTopic, account.ApproveQueueName)
	defer func(amqpChannel *amqp.Channel) {
		err := amqpChannel.Close()
		if err != nil {

		}
	}(amqpChannel)
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	stopChan := make(chan bool)
	go func() {
		for d := range messageChannel {
			fmt.Println(string(d.Body))
			approve := &account.AccountVerifyApprRequest{}
			err := json.Unmarshal(d.Body, approve)
			if err != nil {
				fmt.Println(err)
			}
			res := auth.VerifyUser(approve.UserId, auth.RegistrationStatusApproved)
			t, err := json.Marshal(&account.AccountVerifyApprResponse{Yes: res})
			if len(d.ReplyTo) != 0 {
				err := amqpChannel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					Body:          t,
					CorrelationId: d.CorrelationId,
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
	<-stopChan
}
