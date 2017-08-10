package log
import (
	"log"
	"os"
)

var (
	logger *log.Logger
)

func NewLog() *log.Logger {
	if logger == nil {
		logFileName := "go.log"
		logFile, err := os.Create(logFileName)
		if err != nil {
			log.Fatal("create logfile error")
		}
		logger = log.New(logFile, "[Error]", log.Llongfile)
	}
	return logger
}
