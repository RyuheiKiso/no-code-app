package quic

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/quic-go/quic-go"
)

type Client struct {
	// QUICセッション
	session quic.Connection
	// 接続先アドレス
	address string
	// 再試行回数
	retryAttempts int
	// 再試行間隔
	retryDelay time.Duration
}

// 新しいクライアントを作成する関数
// address: 接続先アドレス
// retryAttempts: 接続の再試行回数
// retryDelay: 再試行間隔
func NewClient(address string, retryAttempts int, retryDelay time.Duration) (*Client, error) {
	// クライアント構造体を初期化
	client := &Client{
		address:       address,
		retryAttempts: retryAttempts,
		retryDelay:    retryDelay,
	}
	// 接続を試みる
	err := client.connect()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// 接続を確立する関数
func (c *Client) connect() error {
	var session quic.Connection
	var err error

	// 再試行回数に基づいて接続を試みる
	for i := 0; i < c.retryAttempts; i++ {
		// タイムアウト付きのコンテキストを作成
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// TLS設定を作成
		tlsConfig := &tls.Config{InsecureSkipVerify: true}
		// QUIC設定を作成
		quicConfig := &quic.Config{}
		// 指定されたアドレスに接続を試みる
		session, err = quic.DialAddr(ctx, c.address, tlsConfig, quicConfig)
		if err == nil {
			// 接続が成功した場合
			log.Printf("Successfully connected to %s", c.address)
			c.session = session
			return nil
		}
		// 接続が失敗した場合
		log.Printf("Failed to dial address %s: %v (attempt %d/%d)", c.address, err, i+1, c.retryAttempts)
		// 再試行前に指定された間隔だけ待機
		time.Sleep(c.retryDelay)
	}

	return err
}

// セッションを閉じる関数
func (c *Client) Close() {
	// セッションをエラーなしで閉じる
	if err := c.session.CloseWithError(0, ""); err != nil {
		log.Fatalf("Failed to close session: %v", err)
	}
	log.Println("Session closed successfully")
}

// メッセージを送信する関数
// ctx: コンテキスト
// message: 送信するメッセージ
func (c *Client) SendMessage(ctx context.Context, message []byte) ([]byte, error) {
	// 接続されていない場合は再接続を試みる
	if !c.IsConnected() {
		if err := c.connect(); err != nil {
			return nil, fmt.Errorf("not connected and failed to reconnect: %v", err)
		}
	}

	// ストリームを同期的に開く
	stream, err := c.session.OpenStreamSync(ctx)
	if err != nil {
		log.Printf("Failed to open stream: %v", err)
		return nil, err
	}
	defer stream.Close()

	// メッセージを送信
	_, err = stream.Write(message)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return nil, err
	}

	// レスポンスを読み取る
	response := make([]byte, 1024)
	n, err := stream.Read(response)
	if err != nil {
		log.Printf("Failed to read response: %v", err)
		return nil, err
	}

	return response[:n], nil
}

// メッセージを非同期に送信する関数
// ctx: コンテキスト
// message: 送信するメッセージ
// responseChan: レスポンスを受け取るチャネル
// errorChan: エラーを受け取るチャネル
func (c *Client) SendMessageAsync(ctx context.Context, message []byte, responseChan chan<- []byte, errorChan chan<- error) {
	go func() {
		// メッセージを送信
		response, err := c.SendMessage(ctx, message)
		if err != nil {
			errorChan <- err
			return
		}
		// レスポンスをチャネルに送信
		responseChan <- response
	}()
}

// メッセージを受信する関数
// ctx: コンテキスト
func (c *Client) ReceiveMessage(ctx context.Context) ([]byte, error) {
	// 接続されていない場合はエラーを返す
	if !c.IsConnected() {
		return nil, fmt.Errorf("not connected")
	}

	// ストリームを受け入れる
	stream, err := c.session.AcceptStream(ctx)
	if err != nil {
		log.Printf("Failed to accept stream: %v", err)
		return nil, err
	}
	defer stream.Close()

	// メッセージを読み取る
	response := make([]byte, 1024)
	n, err := stream.Read(response)
	if err != nil {
		log.Printf("Failed to read message: %v", err)
		return nil, err
	}

	return response[:n], nil
}

// 接続されているかを確認する関数
func (c *Client) IsConnected() bool {
	return c.session != nil && c.session.Context().Err() == nil
}

// 接続状態を取得する関数
func (c *Client) GetConnectionState() quic.ConnectionState {
	return c.session.ConnectionState()
}

// 再試行回数を設定する関数
// attempts: 新しい再試行回数
func (c *Client) SetRetryAttempts(attempts int) {
	c.retryAttempts = attempts
	log.Printf("Retry attempts set to %d", attempts)
}

// 再試行間隔を設定する関数
// delay: 新しい再試行間隔
func (c *Client) SetRetryDelay(delay time.Duration) {
	c.retryDelay = delay
	log.Printf("Retry delay set to %v", delay)
}

// 接続状態をログに出力する関数
func (c *Client) LogConnectionState() {
	state := c.GetConnectionState()
	log.Printf("Connection state: %+v", state)
}
