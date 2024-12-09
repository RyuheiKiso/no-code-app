package logger

import (
	"log"
	"os"
)

// ロガーの初期化
func InitLogger() *log.Logger {
	return log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
