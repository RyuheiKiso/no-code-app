package config

import (
	"log"
	"os"
)

// 環境変数を取得する関数
func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("環境変数 %s が見つかりません。デフォルト値 %s を使用します。", key, defaultValue)
		return defaultValue
	}
	return value
}
