package httpclient

import (
	"net/http"
	"time"
)

// HTTPクライアントの初期化
func NewHTTPClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}
