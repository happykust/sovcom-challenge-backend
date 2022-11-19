package app

import (
	"currency-parser/internal/app/parsers/cryptocurrency"
	"currency-parser/internal/app/parsers/currency"
)

func StartParsers() {
	go cryptocurrency.Setup()
	go currency.Setup()
	select {}
}
