package amqp_easier

import (
	"context"
	logger "deals/pkg/logging"
	LoggerTypes "deals/pkg/logging/types"
	"fmt"
	"math/rand"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishConstructor(connName string, exchangeName string, exchangeType string, routingKey *string, body *[]byte) []byte {
	message := make(chan []byte)

	req, err := publish(connName, message, exchangeName, exchangeType, *routingKey, *body)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Publish err", err)
	}

	return req
}

func publish(connName string, message chan []byte, exchange, exchangeType, routingKey string, body []byte) ([]byte, error) {

	channel, connection := ConnectToRabbit(connName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	defer connection.Close()
	defer channel.Close()

	logger.Log(LoggerTypes.INFO, "RabbitMQ exchange", nil)

	if err := channel.ExchangeDeclare(exchange, exchangeType, true, false, false, false, nil); err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Exchange Declare: ", err)
	}

	replyQueue, err := channel.QueueDeclare("", false, false, true, false, nil)

	Id := randomId()

	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "RabbitMQ QueueDeclare", err)
	}

	if err = channel.PublishWithContext(ctx,
		exchange, routingKey, false, false,
		amqp.Publishing{Headers: amqp.Table{}, ContentType: "text/plain", ContentEncoding: "", ReplyTo: replyQueue.Name, CorrelationId: Id, Body: body, DeliveryMode: amqp.Transient, Priority: 0}); err != nil {
		logger.Log(LoggerTypes.CRITICAL, "RabbitMQ publish", err)

	}

	messageChannel, err := channel.Consume(replyQueue.Name, "", true, false, false, false, nil)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "RabbitMQ consume", err)
	}

	go waitMessage(messageChannel, message, Id)
	select {
	case x := <-message:
		fmt.Println("ok")
		fmt.Println(string(x))
		return x, nil
	case <-ctx.Done():
		fmt.Println("timeout")
		return nil, fmt.Errorf("timeout")
	}
}

func waitMessage(messageChannel <-chan amqp.Delivery, message chan []byte, Id string) {
	fmt.Println("waitMessage")
	for msg := range messageChannel {
		fmt.Println(Id)
		fmt.Println(msg.CorrelationId)
		if msg.CorrelationId == Id {
			message <- msg.Body
			return
		}
		logger.Log(LoggerTypes.CRITICAL, "CorrelationId not equal", nil)
	}
}

func randomId() string {
	rand.Seed(time.Now().UnixNano())
	block1 := rand.Intn(999999-100000) + 100000
	block2 := rand.Intn(999999-100000) + 100000
	block3 := rand.Intn(999999-100000) + 100000
	block4 := rand.Intn(999999-100000) + 100000
	return fmt.Sprintf("%d-%d-%d-%d", block1, block2, block3, block4)
}
