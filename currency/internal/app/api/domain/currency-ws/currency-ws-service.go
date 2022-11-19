package currency_ws

import (
	"currency/internal/app/api/types"
	logger "currency/pkg/logging"
	LoggerTypes "currency/pkg/logging/types"
	"encoding/json"
	"github.com/gorilla/websocket"
)

func SendWSResponse(conn *websocket.Conn, t int, status bool, message string, data interface{}) error {
	response := types.WebSocketOutgoingMessage{
		Status:  status,
		Message: message,
		Data:    data,
	}
	responseMarshaled, _ := json.Marshal(response)
	err := conn.WriteMessage(t, responseMarshaled)
	if err != nil {
		logger.Log(LoggerTypes.ERROR, "Failed to write message: "+err.Error(), nil)
	}
	return err
}
