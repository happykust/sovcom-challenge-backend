package support

import (
	"github.com/gorilla/websocket"
	"libs/contracts/support/accounts"
	"support/internal/app/api/domain/support/sending/amqp"
	"support/static_vars"
	"support/types"
	ws_data_types "support/types/ws-data-types"
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

		validationResponse := amqp.GetUserDataByAccessToken(accounts.ValidateRequest{AccessToken: authRequest.AccessToken})

		if validationResponse.Status == false {
			_ = SendWSResponse(conn, t, false, static_vars.InvalidToken, nil)
			return false, false
		}

		authorizedConnections[conn] = true
		authorizedUsersDatas[conn] = types.UserData{
			UserID:    validationResponse.UserID,
			FirstName: validationResponse.FirstName,
			LastName:  validationResponse.LastName,
			Role:      types.Role(validationResponse.Role),
		}

		_ = SendWSResponse(conn, t, true, static_vars.Authorized, nil)

		return true, true
	}

	// If user not authorized and not sent auth request, return false
	_ = SendWSResponse(conn, t, false, static_vars.Unauthorized, nil)
	return false, false
}
