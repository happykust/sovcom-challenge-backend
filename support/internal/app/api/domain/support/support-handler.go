package support

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	logger "support/pkg/logging"
	LoggerTypes "support/pkg/logging/types"
	"support/static_vars"
	"support/types"
)

// Upgrade simple HTTP connection to WebSockets connection
var UpgraderHTTPToWebSockets = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Variable to store all connections
var allConnections = make(map[*websocket.Conn]bool)

// Variable to store all connections and their authorization status
var authorizedConnections = make(map[*websocket.Conn]bool)

// Variable to store all connections and assigned user data
var authorizedUsersDatas = make(map[*websocket.Conn]types.UserData)

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log(LoggerTypes.INFO, "[WebsocketHandler] New connection established from "+r.RemoteAddr, nil)
	conn, err := UpgraderHTTPToWebSockets.Upgrade(w, r, nil)
	if err != nil {
		logger.Log(LoggerTypes.ERROR, "Failed to set websocket upgrade: "+err.Error(), nil)
		return
	}

	allConnections[conn] = true
	authorizedConnections[conn] = false

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var message types.WebSocketIncomingMessage
		unmarshalingErr := json.Unmarshal(msg, &message)

		if unmarshalingErr != nil {
			_ = SendWSResponse(conn, t, false, static_vars.InvalidRequest, nil)
			continue
		}

		authorized, wasAuthorizedNow := AuthorizeRequestMiddleware(t, conn, message)
		if authorized == false || wasAuthorizedNow == true {
			continue
		}

		switch message.Action {
		case "create-ticket":
			CreateTicketActionHandler(conn, t)
		case "send-message":
			SendMessageActionHandler(conn, t, message)
		case "close-ticket":
			CloseTicketActionHandler(conn, t, message)
		case "get-tickets":
			GetTicketsActionHandler(conn, t)
		case "get-ticket-messages":
			GetTicketMessagesActionHandler(conn, t, message)
		default:
			UnhandledActionHandler(conn, t)
		}
	}
}
