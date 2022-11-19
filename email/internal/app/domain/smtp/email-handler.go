package smtp

import (
	"email/internal/app/domain/smtp/sending/amqp"
	logger "email/pkg/logging"
	LoggerTypes "email/pkg/logging/types"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	emailAccounts "libs/contracts/email/accounts"
	"log"
	"os"
	"strings"
)

func SendEmail(email string, subject string, body string) string {
	smtpHost, smtpUser, smtpPass, smtpPort := os.Getenv("SMTP_HOST"), os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"), os.Getenv("SMTP_PORT")
	auth := sasl.NewPlainClient("", smtpUser, smtpPass)

	to := []string{email}
	msg := strings.NewReader("To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		"" + body + "\r\n")
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, to, msg)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "ERROR_FAILED_TO_SEND_EMAIL", err)
	}

	log.Println("Mail sent successfully")
	return "Mail sent successfully"
}

func SendEmailByUserID(subject string, body string, user_id uint) string {
	userEmail := amqp.GetEmailByUserIDFromAccounts(emailAccounts.SearchEmailRequest{UserID: user_id})
	return SendEmail(userEmail.Email, subject, body)
}
