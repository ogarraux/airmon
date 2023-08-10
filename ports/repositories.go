package ports

import "github.com/ogarraux/airmon/domain"

type MeasurementStoragePort interface {
	StoreMeasurement(domain.AirMeasurement) error
}

type ConfigurationPort interface {
	GetDatastores() []domain.DataStore
	GetSensors() []domain.Sensor
	GetInterval() uint
	GetMonitorEndpoint() string
}

type SensorPort interface {
	CollectMeasurements() (domain.AdditionalMeasurement, error)
}
