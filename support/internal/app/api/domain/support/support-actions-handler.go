package support

import (
	"github.com/gorilla/websocket"
	"strings"
	"support/static_vars"
	"support/types"
	ws_data_types "support/types/ws-data-types"
)

func CreateTicketActionHandler(conn *websocket.Conn, t int) {
	userData := authorizedUsersDatas[conn]

	if userData.Role == types.RoleAdmin {
		_ = SendWSResponse(conn, t, false, static_vars.AccessDenied, nil)
		return
	}

	ticket := GetTicketByUserIDRepository(authorizedUsersDatas[conn].UserID)
	if ticket.ID != 0 {
		_ = SendWSResponse(conn, t, false, static_vars.TicketAlreadyExist, nil)
		return
	}

	createdTicket := CreateTicketRepository(Ticket{
		UserID:    userData.UserID,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
	})

	_ = SendWSResponse(conn, t, true, static_vars.TicketCreated, createdTicket)
}

func SendMessageActionHandler(conn *websocket.Conn, t int, message types.WebSocketIncomingMessage) {
	userData := authorizedUsersDatas[conn]

	if userData.Role == types.RoleAdmin {
		SendMessageAdminActionHandler(conn, t, message, userData)
		return
	}
	SendMessageUserActionHandler(conn, t, message, userData)
}

func SendMessageUserActionHandler(conn *websocket.Conn, t int, message types.WebSocketIncomingMessage,
	userData types.UserData) {

	sendMessageRequest := ws_data_types.SendMessageUserRequest{Message: message.Data.Message}

	noSpaceMessage := strings.TrimSpace(sendMessageRequest.Message)

	if noSpaceMessage == "" {
		_ = SendWSResponse(conn, t, false, static_vars.EmptyMessage, nil)
		return
	}

	ticket := GetTicketByUserIDRepository(userData.UserID)
	if ticket.ID == 0 {
		_ = SendWSResponse(conn, t, false, static_vars.TicketNotExist, nil)
		return
	}

	createdMessage := CreateTicketMessageRepository(TicketMessage{
		Ticket:  ticket,
		FromID:  userData.UserID,
		Message: message.Data.Message,
		Role:    Role(userData.Role),
	})

	_ = SendWSResponse(conn, t, true, static_vars.MessageSent, createdMessage)
}

func SendMessageAdminActionHandler(conn *websocket.Conn, t int, message types.WebSocketIncomingMessage,
	userData types.UserData) {
	sendMessageAdminRequest := ws_data_types.SendMessageAdminRequest{
		TicketID: message.Data.TicketID,
		Message:  message.Data.Message,
	}

	noSpaceMessage := strings.TrimSpace(sendMessageAdminRequest.Message)

	if noSpaceMessage == "" {
		_ = SendWSResponse(conn, t, false, static_vars.EmptyMessage, nil)
		return
	}

	ticket := GetTicketByIDRepository(sendMessageAdminRequest.TicketID)
	if ticket.ID == 0 {
		_ = SendWSResponse(conn, t, false, static_vars.TicketNotExist, nil)
		return
	}

	createdMessage := CreateTicketMessageRepository(TicketMessage{
		Ticket:  ticket,
		FromID:  userData.UserID,
		Message: sendMessageAdminRequest.Message,
		Role:    Role(userData.Role),
	})

	_ = SendWSResponse(conn, t, true, static_vars.MessageSent, createdMessage)
}

func GetTicketsActionHandler(conn *websocket.Conn, t int) {
	userData := authorizedUsersDatas[conn]

	if userData.Role == types.RoleUser {
		_ = SendWSResponse(conn, t, false, static_vars.AccessDenied, nil)
		return
	}

	tickets := GetAllTicketsRepository()

	_ = SendWSResponse(conn, t, true, static_vars.Empty, tickets)
}

func CloseTicketActionHandler(conn *websocket.Conn, t int, message types.WebSocketIncomingMessage) {
	userData := authorizedUsersDatas[conn]

	if userData.Role == types.RoleAdmin {
		CloseTicketAdminActionHandler(conn, t, message)
		return
	}

	CloseTicketUserActionHandler(conn, t, userData)
}

func CloseTicketUserActionHandler(conn *websocket.Conn, t int, userData types.UserData) {

	ticket := GetTicketByUserIDRepository(userData.UserID)
	if ticket.ID == 0 {
		_ = SendWSResponse(conn, t, false, static_vars.TicketNotExist, nil)
		return
	}

	CloseTicketRepository(ticket.ID)
	_ = SendWSResponse(conn, t, true, static_vars.TicketClosed, nil)
}

func CloseTicketAdminActionHandler(conn *websocket.Conn, t int, message types.WebSocketIncomingMessage) {
	closeTicketAdminRequest := ws_data_types.CloseTicketRequest{TicketID: message.Data.TicketID}

	ticket := GetTicketByIDRepository(closeTicketAdminRequest.TicketID)
	if ticket.ID == 0 {
		_ = SendWSResponse(conn, t, false, static_vars.TicketNotExist, nil)
		return
	}

	CloseTicketRepository(ticket.ID)
	_ = SendWSResponse(conn, t, true, static_vars.TicketClosed, nil)
}

func GetTicketMessagesActionHandler(conn *websocket.Conn, t int, message types.WebSocketIncomingMessage) {
	userData := authorizedUsersDatas[conn]

	if userData.Role == types.RoleUser {
		_ = SendWSResponse(conn, t, false, static_vars.AccessDenied, nil)
		return
	}

	getTicketMessagesRequest := ws_data_types.GetTicketMessagesRequest{TicketID: message.Data.TicketID}

	ticket := GetTicketByIDRepository(getTicketMessagesRequest.TicketID)
	if ticket.ID == 0 {
		_ = SendWSResponse(conn, t, false, static_vars.TicketNotExist, nil)
		return
	}

	messages := GetTicketMessagesByTicketIDRepository(ticket.ID)

	_ = SendWSResponse(conn, t, true, static_vars.Empty, messages)
}

func UnhandledActionHandler(conn *websocket.Conn, t int) {
	_ = SendWSResponse(conn, t, false, static_vars.ActionNotFound, nil)
}
