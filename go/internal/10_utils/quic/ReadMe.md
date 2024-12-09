# QUIC Client

このリポジトリには、QUICプロトコルを使用して通信を行うためのクライアント実装が含まれています。`client.go`ファイルには、QUICセッションの管理、メッセージの送受信、および接続の再試行ロジックが実装されています。

## 機能

- QUICセッションの確立および管理
- メッセージの同期および非同期送信
- メッセージの受信
- 接続の再試行ロジック
- 接続状態の取得およびログ出力

## 使用方法

### クライアントの作成

新しいクライアントを作成するには、`NewClient`関数を使用します。この関数は接続先アドレス、再試行回数、および再試行間隔を引数として受け取ります。

```go
client, err := quic.NewClient("example.com:443", 3, 2*time.Second)
if err != nil {
    log.Fatalf("Failed to create client: %v", err)
}
```

### メッセージの送信

メッセージを送信するには、`SendMessage`関数を使用します。この関数はコンテキストと送信するメッセージを引数として受け取ります。

```go
response, err := client.SendMessage(context.Background(), []byte("Hello, QUIC!"))
if err != nil {
    log.Fatalf("Failed to send message: %v", err)
}
fmt.Printf("Received response: %s\n", response)
```

### メッセージの非同期送信

メッセージを非同期に送信するには、`SendMessageAsync`関数を使用します。この関数はコンテキスト、送信するメッセージ、レスポンスを受け取るチャネル、およびエラーを受け取るチャネルを引数として受け取ります。

```go
responseChan := make(chan []byte)
errorChan := make(chan error)
client.SendMessageAsync(context.Background(), []byte("Hello, QUIC!"), responseChan, errorChan)

select {
case response := <-responseChan:
    fmt.Printf("Received response: %s\n", response)
case err := <-errorChan:
    log.Fatalf("Failed to send message: %v", err)
}
```

### メッセージの受信

メッセージを受信するには、`ReceiveMessage`関数を使用します。この関数はコンテキストを引数として受け取ります。

```go
message, err := client.ReceiveMessage(context.Background())
if err != nil {
    log.Fatalf("Failed to receive message: %v", err)
}
fmt.Printf("Received message: %s\n", message)
```

### 接続状態の取得およびログ出力

接続状態を取得するには、`GetConnectionState`関数を使用します。また、接続状態をログに出力するには、`LogConnectionState`関数を使用します。

```go
state := client.GetConnectionState()
fmt.Printf("Connection state: %+v\n", state)

client.LogConnectionState()
```

### 再試行回数および再試行間隔の設定

再試行回数および再試行間隔を動的に設定するには、`SetRetryAttempts`および`SetRetryDelay`関数を使用します。

```go
client.SetRetryAttempts(5)
client.SetRetryDelay(1 * time.Second)
```

### セッションの閉鎖

セッションを閉じるには、`Close`関数を使用します。

```go
client.Close()
```
