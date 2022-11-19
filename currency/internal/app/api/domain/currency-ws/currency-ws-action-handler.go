package currency_ws

import (
	"currency/internal/app/api/static_vars"
	"currency/internal/app/api/types"
	"currency/pkg/core/database"
	"github.com/gorilla/websocket"
	"strings"
)

func SubscribeToCurrency(conn *websocket.Conn, t int, message types.WebSocketIncomingMessage) {
	noSpaceSubscribeTo := strings.TrimSpace(message.Data.SubscribeTo)
	if noSpaceSubscribeTo == "" {
		_ = SendWSResponse(conn, t, false, static_vars.InvalidSubscribeTo, nil)
		return
	}

	userData := authorizedUsersDatas[conn]
	userData.SubscribeTo = message.Data.SubscribeTo

	authorizedUsersDatas[conn] = userData
	_ = SendWSResponse(conn, t, true, static_vars.Subscribed, nil)
}

func UnsubscribeFromCurrency(conn *websocket.Conn, t int) {
	userData := authorizedUsersDatas[conn]
	userData.SubscribeTo = ""

	authorizedUsersDatas[conn] = userData
	_ = SendWSResponse(conn, t, true, static_vars.Unsubscribed, nil)
}

func GetTickerGroupsHandler(conn *websocket.Conn, t int) {
	tickersGroups := database.Redis.SMembers("tickers-groups").Val()
	_ = SendWSResponse(conn, t, true, static_vars.Empty, tickersGroups)
}

func UnhandledActionHandler(conn *websocket.Conn, t int) {
	_ = SendWSResponse(conn, t, false, static_vars.ActionNotFound, nil)
}
