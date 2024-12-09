package errorhandler

import (
	"log"
)

// エラーハンドリング関数
func HandleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
