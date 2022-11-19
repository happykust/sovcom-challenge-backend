package LoggerTypes

type ErrorLevel string

const (
	CRITICAL ErrorLevel = "CRITICAL"
	ERROR               = "ERROR"
	WARNING             = "WARNING"
	INFO                = "INFO"
	DEBUG               = "DEBUG"
)

type Log struct {
	ID      string
	TYPE    ErrorLevel
	MESSAGE string
	ERROR   error
}
