package currency_ws

import (
	"currency/internal/app/api/static_vars"
	"currency/internal/app/api/types"
	logger "currency/pkg/logging"
	LoggerTypes "currency/pkg/logging/types"
	"encoding/json"
	"github.com/gorilla/websocket"
	"libs/contracts/currency"
	"net/http"
)

// Upgrade simple HTTP connection to WebSockets connection
var UpgraderHTTPToWebSockets = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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
		case "subscribe":
			SubscribeToCurrency(conn, t, message)
		case "unsubscribe":
			UnsubscribeFromCurrency(conn, t)
		case "get-tickers-groups":
			GetTickerGroupsHandler(conn, t)
		default:
			UnhandledActionHandler(conn, t)
		}
	}
}

func ReceiveCurrencyUpdate(currencyUpdate currency.CurrencyUpdateRequestToCurrency) {
	for conn := range authorizedConnections {
		userData := authorizedUsersDatas[conn]
		if userData.SubscribeTo == currencyUpdate.TickerGroup {
			_ = SendWSResponse(conn, websocket.TextMessage, true, static_vars.TickerGroupUpdated,
				currencyUpdate.Data)
		}
	}
}
