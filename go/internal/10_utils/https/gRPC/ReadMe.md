# gRPC クライアントモジュール

このディレクトリには、gRPC クライアント接続を作成および管理するための共通モジュールが含まれている。

## ファイル構成

- `client.go`: gRPC クライアント接続を作成および管理するための主要なコードが含まれている。

## 使用方法

### クライアント接続の作成

`NewClientConn` 関数を使って、新しい gRPC クライアント接続を作成する。

以下は、基本的な使用例：

```go
import (
    "log"
    "time"
    "go/internal/10_utils/https/gRPC"
)

func main() {
    // gRPCサーバーのアドレス
    address := "localhost:50051"
    // 接続のタイムアウト時間
    timeout := 5 * time.Second
    // 最大リトライ回数
    maxRetries := 3
    // リトライのバックオフ時間
    retryBackoff := 2 * time.Second
    // 認証トークン
    authToken := "your-auth-token"

    // 新しいgRPCクライアント接続を作成
    conn, err := grpc.NewClientConn(address, timeout, maxRetries, retryBackoff, authToken)
    if err != nil {
        log.Fatalf("接続に失敗しました: %v", err)
    }
    // 関数終了時に接続を閉じる
    defer grpc.CloseClientConn(conn)

    // ここで gRPC クライアントを使ってリクエストを送信する
}
```
