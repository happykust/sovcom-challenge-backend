package types

type WebSocketIncomingMessage struct {
	Action string `json:"action"`
	Data   struct {
		AccessToken string `json:"access_token" required:"false"`
		SubscribeTo string `json:"subscribe_to" required:"false"`
	} `json:"data"`
}

type WebSocketOutgoingMessage struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
