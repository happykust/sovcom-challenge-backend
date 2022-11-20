module loyality

go 1.19

require (
	github.com/google/uuid v1.3.0
	github.com/joho/godotenv v1.4.0
	github.com/rabbitmq/amqp091-go v1.5.0
	gorm.io/driver/postgres v1.4.5
	gorm.io/gorm v1.24.1
	libs v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
)

replace libs => ../libs
