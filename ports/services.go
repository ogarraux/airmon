package ports

// Business Logic Interface

type AirMonitorService interface {
	CollectAllMeasurements() error
	GetConfiguration() ConfigurationPort
}
