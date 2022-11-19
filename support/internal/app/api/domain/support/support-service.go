package support

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	logger "support/pkg/logging"
	LoggerTypes "support/pkg/logging/types"
	"support/types"
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
