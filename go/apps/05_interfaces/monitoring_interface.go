package interfaces

import entities "no-code-app/apps/03_entities"

type MonitoringRepository interface {
	// サービスステータスを取得するメソッド
	GetServiceStatus(serviceName string) (entities.ServiceStatus, error)
}
