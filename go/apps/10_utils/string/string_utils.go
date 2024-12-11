package common

import "strings"

// ToUpperCase は文字列を大文字に変換します
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

// ToLowerCase は文字列を小文字に変換します
func ToLowerCase(s string) string {
	return strings.ToLower(s)
}

// TrimSpaces は文字列の前後の空白を削除します
func TrimSpaces(s string) string {
	return strings.TrimSpace(s)
}
