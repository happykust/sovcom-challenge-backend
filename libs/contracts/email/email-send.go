package email

type Request struct {
	Email   string `json:"email" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

type Response struct {
	Message string `json:"message" binding:"required"`
}
