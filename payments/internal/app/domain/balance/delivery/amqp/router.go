package amqp

//func TestConsumer() {
//
//	messageChannel, amqpChannel, conn := amqp_easier("mem", "priver", "topic", "gdfdfg", "testow")
//
//	defer amqpChannel.Close()
//	defer conn.Close()
//	stopChan := make(chan bool)
//	go func() {
//		fmt.Println("waiting for message")
//		for d := range messageChannel {
//			fmt.Println("waiting for message")
//			// createUser := &account.RegisterRequest{}
//			// err := json.Unmarshal(d.Body, createUser)
//			// if err != nil {
//			//	 logger.Log(LoggerTypes.CRITICAL, constants.ERROR_FAILED_TO_REGISTER_CONSUMER, err)
//			// }
//			// register := balance.SingUp(createUser.Email, createUser.Password, createUser.FullName, createUser.StudentId)
//			// t, err := json.Marshal(&account.RegisterResponse{Message: register.Message, AccessToken: register.AccessToken, RefreshToken: register.RefreshToken})
//			//
//			fmt.Println(string(d.Body))
//			if len(d.ReplyTo) != 0 {
//				ctx := context.Background()
//				msg := amqp.Publishing{
//					ContentType:   "text/plain",
//					Body:          d.Body,
//					CorrelationId: d.CorrelationId,
//				}
//
//				err := amqpChannel.PublishWithContext(ctx, "", "", false, false, msg)
//
//				if err != nil {
//					logger.Log(LoggerTypes.CRITICAL, "Error publishing message", err)
//				}
//			}
//
//		}
//
//	}()
//	<-stopChan
//}
