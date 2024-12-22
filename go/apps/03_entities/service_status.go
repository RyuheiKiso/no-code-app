package entities

import "time"

type ServiceStatus struct {
	// サービス名
	ServiceName string
	// PC名
	PCName string
	// 期間
	FromTo string
	// メモリ使用量
	MemoryUsage float64
	// ディスク使用量
	DiskUsage float64
	// CPU使用量
	CPUUsage float64
	// タイムスタンプ
	Timestamp time.Time
}
