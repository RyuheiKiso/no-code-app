package usecases

import (
	entities "no-code-app/apps/03_entities"
	interfaces "no-code-app/apps/05_interfaces"
)

type MonitoringUseCase interface {
	// サービスステータスを取得するメソッド
	GetServiceStatus(serviceName string) (entities.ServiceStatus, error)
}

type monitoringUseCase struct {
	// リポジトリのインターフェース
	repo interfaces.MonitoringRepository
}

// NewMonitoringUseCaseは、新しいMonitoringUseCaseを初期化します。
func NewMonitoringUseCase(repo interfaces.MonitoringRepository) MonitoringUseCase {
	return &monitoringUseCase{repo: repo}
}

// GetServiceStatusは、指定されたサービスのステータスを取得します。
func (uc *monitoringUseCase) GetServiceStatus(serviceName string) (entities.ServiceStatus, error) {
	return uc.repo.GetServiceStatus(serviceName)
}
