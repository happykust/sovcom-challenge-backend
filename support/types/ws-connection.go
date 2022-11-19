package types

import "github.com/gorilla/websocket"

type WsConnection struct {
	Conn *websocket.Conn
}
