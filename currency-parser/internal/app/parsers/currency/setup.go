package currency

import (
	logger "currency-parser/pkg/logging"
	loggerTypes "currency-parser/pkg/logging/types"
)

func Setup() {
	logger.Log(loggerTypes.INFO, "[Currency-parser | Currency] Start currency parser", nil)

}
