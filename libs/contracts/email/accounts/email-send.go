package emailAccounts

type SearchEmailRequest struct {
	UserID uint `json:"user_id"`
}

type SearchEmailResponse struct {
	Email string `json:"email"`
}
