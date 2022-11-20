package app

import (
	"currency-parser/internal/app/parsers/cryptocurrency"
)

func StartParsers() {
	go cryptocurrency.Setup()
	//go currency.Setup()
	select {}
}
