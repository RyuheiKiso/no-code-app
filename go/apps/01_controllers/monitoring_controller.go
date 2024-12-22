package controllers

import (
	"net/http"
	usecases "no-code-app/apps/02_use_cases"
	entities "no-code-app/apps/03_entities"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// MonitoringControllerは、サービスの監視を行うコントローラーです。
type MonitoringController struct {
	// ユースケースのインターフェース
	useCase usecases.MonitoringUseCase
	// WebSocketのアップグレーダー
	upgrader websocket.Upgrader
	// 接続されているクライアントのマップ
	clients map[*websocket.Conn]bool
	// サービスステータスをブロードキャストするためのチャネル
	broadcast chan entities.ServiceStatus
}

// NewMonitoringControllerは、新しいMonitoringControllerを初期化します。
func NewMonitoringController(router *gin.Engine, useCase usecases.MonitoringUseCase) {
	controller := &MonitoringController{
		useCase: useCase,
		upgrader: websocket.Upgrader{
			// 読み取りバッファサイズ
			ReadBufferSize: 1024,
			// 書き込みバッファサイズ
			WriteBufferSize: 1024,
		},
		// クライアントマップの初期化
		clients: make(map[*websocket.Conn]bool),
		// ブロードキャストチャネルの初期化
		broadcast: make(chan entities.ServiceStatus),
	}
	// サービスステータスを取得するエンドポイント
	router.GET("/monitor/:serviceName", controller.GetServiceStatus)
	// WebSocketハンドラーのエンドポイント
	router.GET("/ws", controller.WebSocketHandler)
	// メッセージハンドリングのゴルーチンを開始
	go controller.handleMessages()
}

// GetServiceStatusは、指定されたサービスのステータスを取得します。
func (ctrl *MonitoringController) GetServiceStatus(c *gin.Context) {
	// パスパラメータからサービス名を取得
	serviceName := c.Param("serviceName")
	// サービスステータスを取得
	status, err := ctrl.useCase.GetServiceStatus(serviceName)
	if err != nil {
		// エラーレスポンスを返す
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// ステータスをJSON形式で返す
	c.JSON(http.StatusOK, status)
}

// WebSocketHandlerは、WebSocket接続を処理します。
func (ctrl *MonitoringController) WebSocketHandler(c *gin.Context) {
	// WebSocket接続をアップグレード
	conn, err := ctrl.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// エラーレスポンスを返す
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 接続終了時にクローズ
	defer conn.Close()
	// クライアントをマップに追加
	ctrl.clients[conn] = true

	for {
		// メッセージを読み取る
		_, _, err := conn.ReadMessage()
		if err != nil {
			// エラーが発生した場合、クライアントをマップから削除
			delete(ctrl.clients, conn)
			break
		}
	}
}

// handleMessagesは、ブロードキャストメッセージを処理します。
func (ctrl *MonitoringController) handleMessages() {
	for {
		// ブロードキャストチャネルからメッセージを受信
		msg := <-ctrl.broadcast
		for client := range ctrl.clients {
			// クライアントにメッセージを送信
			err := client.WriteJSON(msg)
			if err != nil {
				// エラーが発生した場合、クライアントをクローズ
				client.Close()
				// クライアントをマップから削除
				delete(ctrl.clients, client)
			}
		}
	}
}
