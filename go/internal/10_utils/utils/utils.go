package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// ランダムな文字列を生成する関数
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
