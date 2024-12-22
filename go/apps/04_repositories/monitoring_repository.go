package repositories

import (
	entities "no-code-app/apps/03_entities"
	interfaces "no-code-app/apps/05_interfaces"
	"no-code-app/apps/10_utils/quic"
)

var _ interfaces.MonitoringRepository = (*MonitoringRepository)(nil)

type MonitoringRepository struct {
	// QUICクライアント
	quicClient *quic.Client
}

// GetServiceStatusは、指定されたサービスのステータスを取得します。
func (m *MonitoringRepository) GetServiceStatus(serviceName string) (entities.ServiceStatus, error) {
	panic("unimplemented")
}

// NewMonitoringRepositoryは、新しいMonitoringRepositoryを初期化します。
func NewMonitoringRepository(client *quic.Client) *MonitoringRepository {
	return &MonitoringRepository{quicClient: client}
}
