module email

go 1.19

require (
	github.com/emersion/go-sasl v0.0.0-20220912192320-0145f2c60ead
	github.com/emersion/go-smtp v0.15.0
	github.com/joho/godotenv v1.4.0
	github.com/rabbitmq/amqp091-go v1.5.0
	libs v0.0.0-00010101000000-000000000000
)

replace libs => ../libs
