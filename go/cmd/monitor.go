package main

import (
	"log"
	controllers "no-code-app/apps/01_controllers"
	usecases "no-code-app/apps/02_use_cases"
	repositories "no-code-app/apps/04_repositories"
	"no-code-app/apps/10_utils/quic"
	"time"

	"github.com/gin-gonic/gin"
)

func monitoring() {
	router := gin.Default()

	// QUICクライアントを作成
	quicClient, err := quic.NewClient("localhost:443", 3, 2*time.Second)
	if err != nil {
		log.Fatalf("Failed to create QUIC client: %v", err)
	}

	// リポジトリを初期化
	monitoringRepo := repositories.NewMonitoringRepository(quicClient)
	// ユースケースを初期化
	monitoringUseCase := usecases.NewMonitoringUseCase(monitoringRepo)
	// コントローラーを初期化
	controllers.NewMonitoringController(router, monitoringUseCase)

	// サーバーを起動
	router.Run(":8080")
}
