package support

import (
	"support/pkg/core/database"
)

func CreateTicketRepository(ticket Ticket) Ticket {
	database.PG.Create(&ticket)
	// TODO: return created ticket
	return ticket
}

func CreateTicketMessageRepository(ticketMessage TicketMessage) TicketMessage {
	database.PG.Create(&ticketMessage)
	return ticketMessage
}

func CloseTicketRepository(ticketID uint) {
	database.PG.Model(&Ticket{}).Where("id = ?", ticketID).Update("is_open", false)
}

func GetTicketByUserIDRepository(userID uint) Ticket {
	var ticket Ticket
	database.PG.First(&ticket, "user_id = ? AND is_open = ?", userID, true)
	return ticket
}

func GetTicketByIDRepository(ticketID uint) Ticket {
	var ticket Ticket
	database.PG.First(&ticket, "id = ? AND is_open = ?", ticketID, true)
	return ticket
}

func GetAllTicketsRepository() []Ticket {
	var tickets []Ticket
	database.PG.Find(&tickets)
	return tickets
}

func GetTicketMessagesByTicketIDRepository(ticketID uint) []TicketMessage {
	var ticketMessages []TicketMessage
	database.PG.Find(&ticketMessages, "ticket_id = ?", ticketID)
	return ticketMessages
}
