package ws_data_types

type SendMessageUserRequest struct {
	Message string `json:"message"`
}

type SendMessageAdminRequest struct {
	TicketID uint   `json:"ticket_id"`
	Message  string `json:"message"`
}
