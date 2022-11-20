package logger

import (
	LoggerTypes "deals/pkg/logging/types"
	"fmt"
	"math/rand"
	"time"
)

func randomId() string {
	rand.Seed(time.Now().UnixNano())
	block1 := rand.Intn(999999-100000) + 100000
	block2 := rand.Intn(999999-100000) + 100000
	block3 := rand.Intn(999999-100000) + 100000
	block4 := rand.Intn(999999-100000) + 100000
	return fmt.Sprintf("log-%d-%d-%d-%d", block1, block2, block3, block4)
}

func Log(level LoggerTypes.ErrorLevel, message string, err error) {
	id := randomId()
	NewLog := LoggerTypes.Log{
		ID:      id,
		TYPE:    level,
		MESSAGE: message,
		ERROR:   err,
	}
	fmt.Println(NewLog)
}
