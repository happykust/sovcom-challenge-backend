package currency_ws

import (
	"currency/internal/app/api/domain/currency-ws/sending/amqp"
	"currency/internal/app/api/static_vars"
	"currency/internal/app/api/types"
	ws_data_types "currency/internal/app/api/types/ws-data-types"
	"github.com/gorilla/websocket"
	"libs/contracts/currency/currencyToAccounts"
)

type IsAuthenticated bool
type WasAuthenticatedNow bool

func AuthorizeRequestMiddleware(t int, conn *websocket.Conn, request types.WebSocketIncomingMessage) (IsAuthenticated, WasAuthenticatedNow) {
	// Check if user already authorized
	if authorizedConnections[conn] {
		return true, false
	}

	// IF user not authorized, check if he sent auth request
	if request.Action == "auth" {
		authRequest := ws_data_types.AuthRequest{AccessToken: request.Data.AccessToken}

		if authRequest.AccessToken == "" {
			_ = SendWSResponse(conn, t, false, static_vars.InvalidToken, nil)
			return false, false
		}

		validationResponse := amqp.GetUserDataByAccessToken(currencyToAccounts.ValidateRequest{AccessToken: authRequest.AccessToken})

		if validationResponse.Status == false {
			_ = SendWSResponse(conn, t, false, static_vars.InvalidToken, nil)
			return false, false
		}

		authorizedConnections[conn] = true
		authorizedUsersDatas[conn] = types.UserData{
			UserID:      validationResponse.UserID,
			SubscribeTo: "",
		}

		_ = SendWSResponse(conn, t, true, static_vars.Authorized, nil)

		return true, true
	}

	// If user not authorized and not sent auth request, return false
	_ = SendWSResponse(conn, t, false, static_vars.Unauthorized, nil)
	return false, false
}
