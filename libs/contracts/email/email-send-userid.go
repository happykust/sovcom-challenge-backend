package email

type SendEmailByUserIDRequest struct {
	UserID  uint   `json:"user_id"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type SendEmailByUserIDResponse struct {
	Message string `json:"message"`
}
